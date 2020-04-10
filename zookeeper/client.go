package zookeeper

import (
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

type Client struct{
	Conn *zk.Conn
	RootNode string
	Servers []string
}

func NewClient(servers []string, rootNode string, timeout int)(*Client, error){
	client := new(Client)

	conn, _,err := zk.Connect(servers, time.Duration(timeout) * time.Second)
	if err != nil{
		return nil, err
	}
	client.Conn = conn
	client.Servers = servers
	client.RootNode = rootNode

	// check zookeeper root node is exist.
	if err := client.EnsureRootNode(); err!= nil{
		return nil, err
	}
	return client, nil
}

func (c *Client) EnsureRootNode()error{
	exist,_, err := c.Conn.Exists(c.RootNode)
	if err != nil{
		return err
	}
	if !exist {
		_, err := c.Conn.Create(c.RootNode, []byte(""), 0, zk.WorldACL(zk.PermAll))
		if err != nil{
			return err
		}
	}
	return nil
}

