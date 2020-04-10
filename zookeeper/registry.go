package zookeeper

import (
	"encoding/json"
	"github.com/samuel/go-zookeeper/zk"
)

type ServiceNode struct {
	Name string
	Host string
	Port int
}

func (c *Client)Registry(node *ServiceNode)error{
	if err := c.EnsureNode(node.Name); err != nil{
		return nil
	}
	path := c.RootNode + "/" + node.Name + "/"
	data,err := json.Marshal(node)
	if err != nil{
		return nil
	}
	_, err = c.Conn.CreateProtectedEphemeralSequential(path,data,zk.WorldACL(zk.PermAll))
	if err != nil{
		return err
	}
	return nil
}

func (c *Client)EnsureNode(name string)error{
	path := c.RootNode + "/" + name
	exist,_, err := c.Conn.Exists(path)
	if err != nil{
		return nil
	}
	if !exist {
		_, err := c.Conn.Create(path, []byte(""), 0, zk.WorldACL(zk.PermAll))
		if err != nil{
			return err
		}
	}
	return nil
}
