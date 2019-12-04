
server.go

注册新的 ip 到 etcd 的 services/ 中,被 master 监听


master.go

监听 server.go 向 etcd services/ 中监听到的新ip

## 前置工作
需修改 master.go ,runserver/*.go 文件的 ip http://192.168.0.229 为自己etcd 所在 ip

## 启动
启动 etcd 单机集群
```bash
sh start_etcd.sh 1
sh start_etcd.sh 2
sh start_etcd.sh 3
```

启动 master 监听服务
```bash

go  run master.go

```

启动测试添加新的服务器服务,测试启动两台server
```bash
cd runserver
go run server.go
go run server-1.go
```
