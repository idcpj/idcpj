
server.go

注册新的 ip 到 etcd 的 services/ 中,被 master 监听


master.go

监听 server.go 向 etcd services/ 中监听到的新ip

