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
```

## AddInbound
此方法使用 InboundHandlerConfig 配置增加一个入站.

## RemoveInbound
移除一个入站, 使用Tag区分

## AlterInbound
在一个入站代理中添加一个用户(VMess, VLESS, Trojan, ShadowSocks)

