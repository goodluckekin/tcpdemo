/**
 * @Author: EDZ
 * @Description:
 * @File: request
 * @Version: 1.0.0
 * @Date: 2022/7/5 15:10
 */

package znet

import (
	"zinx/ziface"
)

var _ ziface.IRequest = (*Request)(nil)

type Request struct {
	msg        ziface.IMessage
	connection ziface.IConnection
}

func NewRequest(conn ziface.IConnection, msg ziface.IMessage) *Request {
	return &Request{
		msg:        msg,
		connection: conn,
	}
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.connection
}

func (r *Request) GetData() []byte {
	return r.msg.GetMsgData()
}

func (r *Request) GetMsgId() uint32 {
	return r.msg.GetMsgId()
}
