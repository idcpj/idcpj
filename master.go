package main

import (
	"fmt"
	"log"
	"time"


	"github.com/etcd-io/etcd/clientv3"
	"github.com/idcpj/service_discovery/discovery"
)

func main() {

	m, err := discovery.NewMaster([]string{
		"http://192.168.0.229:2379",
		"http://192.168.0.229:22379",
		"http://192.168.0.229:32379",
	}, "services/")

	if err != nil {
		log.Fatal(err)
	}

	m.SetCallBack(func(ev *clientv3.Event,info *discovery.ServiceInfo) {
		log.Printf("[%s] %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value) //[PUT] "services/s-test-1" : "{\"IP\":\"192.168.1.1\"}"
		fmt.Printf("%+v\n", info)

		if ev.Type==clientv3.EventTypePut { // 添加新的
			log.Println("发现新服务...")

		}else if ev.Type==clientv3.EventTypeDelete  { // 删除操作
			//todo
			log.Println("有服务被删除...")
		}

	})


	for {
		log.Printf("nodes num = %d\n",len(m.Nodes))
		time.Sleep(time.Second * 5)
	}

}
