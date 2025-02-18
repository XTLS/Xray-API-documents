package examples

import (
	"context"
	"github.com/xtls/xray-core/app/proxyman/command"
)

/*
获取 Tag 下用户数据, 如下
users:{email:"love@xray.com"  account:{type:"xray.proxy.vmess.Account"  value:"\n$10354ac4-9ec1-4864-ba3e-f5fd35869ef8\x1a\x02\x08\x04"}}
Email 为空时返回所有用户
*/
func getInboundUsers(client command.HandlerServiceClient) (users *command.GetInboundUserResponse, err error) {

	users, err = client.GetInboundUsers(context.Background(), &command.GetInboundUserRequest{
		Tag: "proxy0",
		//Email: "love@xray.com",
	})

	return
}
