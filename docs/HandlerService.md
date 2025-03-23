# HandlerService
> 这部分控制由 HandlerServiceClient 承载

支持的接口方法
```shell
AddInbound()
RemoveInbound()
AlterInbound()
AddOutbound()
RemoveOutbound()
AlterOutbound()
getInboundUsers()
```

## AddInbound
此方法使用 InboundHandlerConfig 配置增加一个入站。

## RemoveInbound
移除一个入站，使用 Tag 区分。

## AlterInbound
在一个入站代理中添加一个用户 (VMess, VLESS, Trojan, ShadowSocks)

## getInboundUsers
获取一个入站代理的用户列表

> Outbound 同理