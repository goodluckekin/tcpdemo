/**
 * @Author: ekin
 * @Description:封包与拆包
 * @File: datapack
 * @Version: 1.0.0
 * @Date: 2022/7/6 9:40
 */

package znet

import (
	"bytes"
	"encoding/binary"
	"zinx/ziface"
)

type DataPack struct{}

func NewDataPack() *DataPack {
	return new(DataPack)
}

func (d DataPack) GetDataPackHeadLen() uint32 {
	//协议id位4字节，长度位4字节
	return 8
}

//封包
func (d DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})

	//写入id
	if err := binary.Write(buf, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	//写入data长度
	if err := binary.Write(buf, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}

	//写入data
	if err := binary.Write(buf, binary.LittleEndian, msg.GetMsgData()); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

//拆包
func (d DataPack) Unpack(b []byte) (ziface.IMessage, error) {
	buf := bytes.NewBuffer(b)
	msg := &Message{}

	//读取id
	if err := binary.Read(buf, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	//读取长度
	if err := binary.Read(buf, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	if msg.GetMsgLen() > 0 {
		//读取数据
		if err := binary.Read(buf, binary.LittleEndian, &msg.Data); err != nil {
			return nil, err
		}
	}

	return msg, nil
}
