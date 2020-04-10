package zookeeper

import (
	"encoding/json"
	"github.com/samuel/go-zookeeper/zk"
)

func (c *Client)GetNode(name string)([]*ServiceNode, error){
	path := c.RootNode + "/" + name

	childes, _, err := c.Conn.Children(path)
	if err != nil{
		if err == zk.ErrNoNode {
			return []*ServiceNode{}, nil
		}
		return nil, err
	}
	var nodes []*ServiceNode
	for _, child := range childes {
		fullPath := path + "/" + child
		data, _, err := c.Conn.Get(fullPath)
		if err != nil{
			if err == zk.ErrNoNode {
				continue
			}
			return nil, err
		}
		node := new(ServiceNode)
		err = json.Unmarshal(data, node)
		if err != nil{
			return nil, err
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}
