/**
 * @Author: EDZ
 * @Description:
 * @File: connection
 * @Version: 1.0.0
 * @Date: 2022/7/5 14:35
 */

package znet

import (
	"fmt"
	"io"
	"net"
	"sync"
	"zinx/ziface"
)

var _ ziface.IConnection = (*Connection)(nil)

type Connection struct {
	connId   uint32
	conn     *net.TCPConn
	isClosed bool
	exitChan chan bool
	handlers ziface.IHandler
	msgChan  chan []byte
	srv      ziface.IServer
	property map[string]interface{}
	pLock    sync.RWMutex
}

func NewConnection(srv ziface.IServer, connId uint32, conn *net.TCPConn, h ziface.IHandler) *Connection {
	return &Connection{
		connId:   connId,
		conn:     conn,
		isClosed: false,
		exitChan: make(chan bool, 1),
		handlers: h,
		msgChan:  make(chan []byte),
		srv:      srv,
		property: make(map[string]interface{}),
	}
}

func (c *Connection) Start() {
	go c.StartReader()
	go c.StartWriter()

	//新增链接
	if err := c.srv.GetConnManager().AddConnection(c); err != nil {
		fmt.Println("connManager add connection error", err)
	}

	//创建连接后的hook
	hook := c.srv.GetStartConnectHook()
	hook(c)
}

//读写分离，读的携程
func (c *Connection) StartReader() {
	for {
		dp := NewDataPack()

		//先获取头部消息
		buf := make([]byte, dp.GetDataPackHeadLen())
		_, err := io.ReadFull(c.conn, buf)
		if err != nil {
			fmt.Printf("【connection】 read head error %v \n", err)
			return
		}

		msg, err := dp.Unpack(buf)
		if err != nil {
			fmt.Printf("【connection】 unpack head error %v \n", err)
			return
		}

		//读取data数据
		dataBuf := make([]byte, msg.GetMsgLen())
		if _, err := io.ReadFull(c.conn, dataBuf); err != nil {
			fmt.Printf("【connection】 read data error %v \n", err)
			return
		}

		msg.SetMsgData(dataBuf)
		fmt.Printf("[connection] recv <- %s \n", string(msg.GetMsgData()))

		//请求放到协程池
		req := NewRequest(c, msg)
		if err := c.srv.Dispatch(req); err != nil {
			fmt.Println("dispatch request error", err)
			return
		}
	}
}

//写的携程
func (c *Connection) StartWriter() {
	for {
		select {
		case buf, ok := <-c.msgChan:
			if ok {
				if _, err := c.conn.Write(buf); err != nil {
					fmt.Printf("conn write msg error %v", err)
					return
				}
			}
		case <-c.exitChan:
			return
		}
	}
}

func (c *Connection) Stop() {
	if c.isClosed {
		return
	}

	//结束链接前的hook
	hook := c.srv.GetStopConnectHook()
	hook(c)

	//移除链接
	if err := c.srv.GetConnManager().DelConnection(c.GetConnID()); err != nil {
		fmt.Println("del coonection err", err)
	}

	defer fmt.Printf("【connection】 closed...")
	c.conn.Close()
	close(c.msgChan)
	close(c.exitChan)
}

func (c *Connection) Send(id uint32, data []byte) error {
	if c.msgChan == nil {
		return fmt.Errorf("write chan is closed")
	}
	dp := NewDataPack()
	msg := NewMessage(id, data)
	buf, err := dp.Pack(msg)
	if err != nil {
		fmt.Printf("消息发送异常：%v", err)
		return err
	}
	c.msgChan <- buf
	return nil
}

func (c *Connection) GetConnID() uint32 {
	return c.connId
}

func (c *Connection) GetConn() *net.TCPConn {
	return c.conn
}

func (c *Connection) SetProperty(key string, val interface{}) {
	c.pLock.Lock()
	defer c.pLock.Unlock()
	c.property[key] = val
}

func (c *Connection) GetProperty(key string) interface{} {
	c.pLock.RLock()
	defer c.pLock.RUnlock()
	if res, exist := c.property[key]; exist {
		return res
	} else {
		return nil
	}
}
