/**
 * @Author: ekin
 * @Description:
 * @File: message
 * @Version: 1.0.0
 * @Date: 2022/7/6 9:28
 */

package znet

import "zinx/ziface"

var _ ziface.IMessage = (*Message)(nil)

type Message struct {
	Id      uint32
	DataLen uint32
	Data    []byte
}

func NewMessage(id uint32, data []byte) *Message {
	return &Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

func (m *Message) GetMsgId() uint32 {
	return m.Id
}

func (m *Message) GetMsgLen() uint32 {
	return m.DataLen
}

func (m *Message) GetMsgData() []byte {
	return m.Data
}

func (m *Message) SetMsgId(id uint32) {
	m.Id = id
}

func (m *Message) SetMsgLen(len uint32) {
	m.DataLen = len
}

func (m *Message) SetMsgData(data []byte) {
	m.Data = data
}
