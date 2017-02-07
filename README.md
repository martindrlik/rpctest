# rpctest
ordinary rpc test with no extra meaning

Example
```
$ server -address 127.0.0.1:8085
Listen(tcp, 127.0.0.1:8085)
Reload called with arguments &{7 8}
```

```
$ client -address 127.0.0.1:8085
Listen(tcp, 127.0.0.1:8085)
Server: 7*8=56
```
