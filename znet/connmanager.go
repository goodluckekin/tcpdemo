/**
 * @Author: ekin
 * @Description: 链接管理
 * @File: connManager
 * @Version: 1.0.0
 * @Date: 2022/7/14 9:28
 */

package znet

import (
	"fmt"
	"sync"
	"zinx/ziface"
)

type ConnManager struct {
	conn map[uint32]ziface.IConnection
	mux  sync.Mutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		conn: make(map[uint32]ziface.IConnection),
	}
}

//新增链接
func (c *ConnManager) AddConnection(connection ziface.IConnection) error {
	c.mux.Lock()
	defer c.mux.Unlock()
	fmt.Println("connManager add connection id:", connection.GetConnID())

	connId := connection.GetConnID()
	if _, exist := c.conn[connId]; exist {
		return nil
	}
	c.conn[connId] = connection
	return nil
}

//获取链接
func (c *ConnManager) GetConnection(id uint32) (ziface.IConnection, error) {
	if connection, exist := c.conn[id]; exist {
		return connection, nil
	}

	return nil, fmt.Errorf("connection not exist id:%d", id)
}

//删除链接
func (c *ConnManager) DelConnection(id uint32) error {
	c.mux.Lock()
	defer c.mux.Unlock()
	fmt.Println("connManager del connection id:", id)
	delete(c.conn, id)
	return nil
}

//清空链接
func (c *ConnManager) ClearConnection() error {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.conn = nil
	return nil
}
