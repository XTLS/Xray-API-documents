package examples

import (
	loggerService "github.com/xtls/xray-core/app/log/command"
	handlerService "github.com/xtls/xray-core/app/proxyman/command"
	routingService "github.com/xtls/xray-core/app/router/command"
	statsService "github.com/xtls/xray-core/app/stats/command"
	"google.golang.org/grpc"
)

// Xray API 监听地址及端口
type BaseConfig struct {
	APIAddress string
	APIPort    uint16
}

type UserInfo struct {
	// For VMess & Trojan
	Uuid string
	// User's Level
	Level uint32
	// Which Inbound will add this user
	InTag string
	// User's Email, it's a unique identifier for users
	Email string
	// For ShadowSocks
	CipherType string
	// For ShadowSocks
	Password string
}

// Xray API 操作
type XrayController struct {
	HsClient handlerService.HandlerServiceClient
	SsClient statsService.StatsServiceClient
	LsClient loggerService.LoggerServiceClient
	RsClient routingService.RoutingServiceClient
	CmdConn  *grpc.ClientConn
}
