/**
提供监听路径path，启动master，当put时加入 map,  delete时 从map去掉
*/
package discovery

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/etcd-io/etcd/clientv3"
)

type  callback  func(ev *clientv3.Event,info *ServiceInfo)


type Master struct {
	Path 		string
	Nodes 		map[string] *Node
	Client 		*clientv3.Client
	callback  	callback


}

// node is a client
type Node struct {
	State	bool
	Key		string
	Info    ServiceInfo
}


func NewMaster(endpoints []string, watchPath string) (*Master,error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:	endpoints,
		DialTimeout: time.Second,
	})

	if err != nil {
		log.Fatal(err)
		return nil,err
	}

	master := &Master {
		Path:	watchPath,
		Nodes:	make(map[string]*Node),
		Client: cli,
	}

	go master.WatchNodes()
	return master,err
}



func (m *Master) AddNode(key string,info *ServiceInfo) {
	node := &Node{
		State:	true,
		Key:	key,
		Info:	*info,
	}

	m.Nodes[node.Key] = node
}
func (m *Master) SetCallBack(call callback) {
	m.callback=call
}
func (m *Master) CallBackPut(ev *clientv3.Event)  *clientv3.Event {
	log.Printf("[%s] %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
	return  ev
}

func (m *Master) CallBackDelete(ev *clientv3.Event) *clientv3.Event {
	log.Printf("[%s] %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
	return  ev
}


func GetServiceInfo(ev *clientv3.Event) *ServiceInfo {
	info := &ServiceInfo{}
	err := json.Unmarshal(ev.Kv.Value, info)
	if err != nil {
		log.Println(err)
	}
	return info
}

func (m *Master) WatchNodes()  {
	rch := m.Client.Watch(context.Background(), m.Path, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case clientv3.EventTypePut:
				fmt.Printf("[%s] %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				// 触发回到函数
				info := GetServiceInfo(ev)
				if m.callback!=nil {
					m.callback(ev,info)
				}
				m.AddNode(string(ev.Kv.Key),info)
			case clientv3.EventTypeDelete:
				fmt.Printf("[%s] %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				if m.callback!=nil {
					m.callback(ev,nil)
				}
				delete(m.Nodes, string(ev.Kv.Key))
			}
		}
	}
}