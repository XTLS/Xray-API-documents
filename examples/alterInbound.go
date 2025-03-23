package examples

import (
	"context"
	"github.com/xtls/xray-core/app/proxyman/command"
	"github.com/xtls/xray-core/common/protocol"
	"github.com/xtls/xray-core/common/serial"
	"github.com/xtls/xray-core/proxy/shadowsocks"
	"github.com/xtls/xray-core/proxy/trojan"
	"github.com/xtls/xray-core/proxy/vmess"
)

// 这部分非常简单, 请先去 structure.go 下查看 UserInfo 结构信息

func addVmessUser(client command.HandlerServiceClient, user *UserInfo) error {
	_, err := client.AlterInbound(context.Background(), &command.AlterInboundRequest{
		// 先确定哪个入站端口将要添加用户
		Tag: user.InTag,
		// 添加用户操作 github.com/xtls/xray-core/app/proxyman/command 中的 AddUserOperation
		Operation: serial.ToTypedMessage(&command.AddUserOperation{
			User: &protocol.User{
				// 用户信息（Level和Email为所有入站用户都需要的信息）
				Level: user.Level,
				Email: user.Email,
				/* 	不同代理类型使用不同的用户信息结构
				请在 github.com/xtls/xray-core/proxy/PROXYTYPE 下寻找 Account 结构体
				*/
				Account: serial.ToTypedMessage(&vmess.Account{
					Id: user.Uuid,
				}),
			},
		}),
	})
	return err
}

func addSSUser(client command.HandlerServiceClient, user *UserInfo) error {
	var ssCipherType shadowsocks.CipherType
	switch user.CipherType {
	case "aes-128-gcm":
		ssCipherType = shadowsocks.CipherType_AES_128_GCM
	case "aes-256-gcm":
		ssCipherType = shadowsocks.CipherType_AES_256_GCM
	case "chacha20-ietf-poly1305":
		ssCipherType = shadowsocks.CipherType_CHACHA20_POLY1305
	}

	_, err := client.AlterInbound(context.Background(), &command.AlterInboundRequest{
		Tag: user.InTag,
		Operation: serial.ToTypedMessage(&command.AddUserOperation{
			User: &protocol.User{
				Level: user.Level,
				Email: user.Email,
				Account: serial.ToTypedMessage(&shadowsocks.Account{
					Password:   user.Password,
					CipherType: ssCipherType,
				}),
			},
		}),
	})
	return err
}

func addTrojanUser(client command.HandlerServiceClient, user *UserInfo) error {
	_, err := client.AlterInbound(context.Background(), &command.AlterInboundRequest{
		Tag: user.InTag,
		Operation: serial.ToTypedMessage(&command.AddUserOperation{
			User: &protocol.User{
				Level: user.Level,
				Email: user.Email,
				Account: serial.ToTypedMessage(&trojan.Account{
					Password: user.Uuid,
				}),
			},
		}),
	})
	return err
}

// Email 为用户唯一标识符, 使用Email配合用户入站Tag来删除用户

func removeUser(client command.HandlerServiceClient, user *UserInfo) error {
	_, err := client.AlterInbound(context.Background(), &command.AlterInboundRequest{
		Tag: user.InTag,
		Operation: serial.ToTypedMessage(&command.RemoveUserOperation{
			Email: user.Email,
		}),
	})
	return err
}
