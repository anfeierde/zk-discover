package main

import (
	"fmt"
	"zk-discover/zookeeper"
)

func main(){
	servers := []string{"127.0.0.1:2181"}
	client, err := zookeeper.NewClient(servers, "/api", 10)
	if err != nil{
		panic(err)
	}

	// registry service
	if err := client.Registry(&zookeeper.ServiceNode{
		Name: "user",
		Host: "127.0.0.1",
		Port: 1001,
	}); err != nil{
		panic(err)
	}
	if err := client.Registry(&zookeeper.ServiceNode{
		Name: "user",
		Host: "127.0.0.1",
		Port: 1002,
	}); err != nil{
		panic(err)
	}


	// service discover.
	serviceNodes , err := client.GetNode("user")
	if err != nil{
		panic(err)
	}
	for _, node := range serviceNodes {
		fmt.Println(node.Name, node.Host, node.Port)
	}
	select {}
}
