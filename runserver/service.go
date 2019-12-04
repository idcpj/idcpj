/**

发现一台新的ip
 */
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/idcpj/service_discovery/discovery"
)

func main() {

	serviceName := "s-test-0"
	serviceInfo := discovery.ServiceInfo{IP:"192.168.1.0",Port:1080}

	s, err := discovery.NewService(serviceName,"services/", serviceInfo,[]string {
		"http://192.168.0.229:2379",
		"http://192.168.0.229:22379",
		"http://192.168.0.229:32379",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name:%s, ip:%s\n", s.Name, s.Info.IP)

	go func() {
		time.Sleep(time.Second*20)
		s.Stop()
	}()

	s.Start()
}