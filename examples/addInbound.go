package examples

import (
	"context"
	"github.com/xtls/xray-core/app/proxyman"
	"github.com/xtls/xray-core/app/proxyman/command"
	"github.com/xtls/xray-core/common/net"
	protocol "github.com/xtls/xray-core/common/protocol"
	"github.com/xtls/xray-core/common/protocol/tls/cert"
	"github.com/xtls/xray-core/common/serial"
	"github.com/xtls/xray-core/core"
	_ "github.com/xtls/xray-core/infra/conf"
	"github.com/xtls/xray-core/proxy/vmess"
	//ssInbound "github.com/xtls/xray-core/proxy/shadowsocks"
	//trojanInbound "github.com/xtls/xray-core/proxy/trojan"
	vmessInbound "github.com/xtls/xray-core/proxy/vmess/inbound"
	"github.com/xtls/xray-core/transport/internet"
	_ "github.com/xtls/xray-core/transport/internet/tcp"
	"github.com/xtls/xray-core/transport/internet/tls"
	"github.com/xtls/xray-core/transport/internet/websocket"
)

func addInbound(client command.HandlerServiceClient) error {

	_, err := client.AddInbound(context.Background(), &command.AddInboundRequest{
		Inbound: &core.InboundHandlerConfig{
			Tag: "proxy0",
			ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
				// 监听端口 12345
				PortList: &net.PortList{
					Range: []*net.PortRange{net.SinglePortRange(12360)},
				},
				// 监听地址, 默认0.0.0.0
				Listen: net.NewIPOrDomain(net.AnyIP),
				// 流量探测
				SniffingSettings: &proxyman.SniffingConfig{
					Enabled:             true,
					DestinationOverride: []string{"http", "tls"},
				},
				// 传输方式
				StreamSettings: &internet.StreamConfig{
					/*
						传输方式名称
						请自行在 github.com/xtls/xray-core/transport/internet/config.pb.go 中寻找支持的协议
						截至 1.3.0 目前支持
						"TCP",
						"UDP",
						"MKCP",
						"WebSocket",
						"HTTP",
						"DomainSocket"
						使用时请一律小写

					*/
					ProtocolName: "websocket",
					TransportSettings: []*internet.TransportConfig{
						{
							ProtocolName: "websocket",
							/*
								选定传输方式后,请去 github.com/xtls/xray-core/transport/internet 下你选定方式的文件夹中导入config结构
								如选定WebSocket则需要使用 github.com/xtls/xray-core/transport/internet/websocket/config.pb.go 中的 Config struct
								结构内容请自行翻看代码（Ctrl + 左键）
							*/
							Settings: serial.ToTypedMessage(&websocket.Config{
								Path: "/web",
								Header: map[string]string{
									"Host": "www.xray.best",
								},
								AcceptProxyProtocol: false,
							},
							),
						},
					},
					/*
						传输层加密
						请在 github.com/xtls/xray-core/transport/internet/ 中选择合适的传输层加密方式
						截至1.3.0 目前支持
						TLS
						XTLS
						留空即为None
					*/
					SecurityType: serial.GetMessageType(&tls.Config{}),
					SecuritySettings: []*serial.TypedMessage{
						serial.ToTypedMessage(&tls.Config{
							//Auto build
							Certificate: []*tls.Certificate{tls.ParseCertificate(cert.MustGenerate(nil))},
						}),
					},
				},
			}),
			/*
				代理设置, 请到 github.com/xtls/xray-core/proxy/ 寻找你想要添加的入站代理类型
				某些类型需要区分 Inbound 与 Outbound 的配置,
				需要区分使用 github.com/xtls/xray-core/proxy/PROXYTYPE/inbound/config.pb.go 中的 Config 结构
				无须区分的使用 github.com/xtls/xray-core/proxy/PROXYTYPE/config.pb.go 的 ServerConfig 结构
			*/
			ProxySettings: serial.ToTypedMessage(&vmessInbound.Config{
				User: []*protocol.User{
					{
						Level: 0,
						Email: "love@xray.com",
						Account: serial.ToTypedMessage(&vmess.Account{
							Id: "10354ac4-9ec1-4864-ba3e-f5fd35869ef8",
						}),
					},
				},
			}),
		},
	})

	return err
}
