/**
 * @Author: EDZ
 * @Description:
 * @File: datapack_test.go
 * @Version: 1.0.0
 * @Date: 2022/7/7 11:32
 */

package znet

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net"
	"testing"
	"time"
)

func TestDataPack_Pack(t *testing.T) {

	//模拟服务端
	listener, err := net.Listen("tcp4", "127.0.0.1:7777")
	assert.Nil(t, err)

	go func() {
		for {
			conn, err := listener.Accept()
			assert.Nil(t, err)

			//读包并拆包
			go func() {
				for {

					datapack := NewDataPack()
					headBuf := make([]byte, datapack.GetDataPackHeadLen())
					_, err := io.ReadFull(conn, headBuf)
					assert.Nil(t, err)

					//读取头部信息
					msg, err := datapack.Unpack(headBuf)
					assert.Nil(t, err)

					//读取数据
					dataBuf := make([]byte, msg.GetMsgLen())
					_, err = io.ReadFull(conn, dataBuf)
					assert.Nil(t, err)
					msg.SetMsgData(dataBuf)

					t.Logf("数据包id:%d 长度:%d 数据:%s", msg.GetMsgId(), msg.GetMsgLen(), string(msg.GetMsgData()))
				}
			}()
		}
	}()

	//模拟客户端
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	assert.Nil(t, err)
	defer conn.Close()

	//模拟发送 packageA
	packageA := &Message{
		Id:      1,
		DataLen: 8,
		Data:    []byte("packageA"),
	}
	packdata := NewDataPack()
	packageabuf, err := packdata.Pack(packageA)
	assert.Nil(t, err)
	_, err = conn.Write(packageabuf)
	assert.Nil(t, err)

	//模拟发送 packageB
	packageB := &Message{
		Id:      2,
		DataLen: 2,
		Data:    []byte("pB"),
	}
	packagebbuf, err := packdata.Pack(packageB)
	assert.Nil(t, err)
	_, err = conn.Write(packagebbuf)
	assert.Nil(t, err)

	//模拟发送 粘包
	packagecbuf := append(packageabuf, packagebbuf...)
	_, err = conn.Write(packagecbuf)
	assert.Nil(t, err)

	time.Sleep(2 * time.Second)
}

func TestDataPack_for(t *testing.T) {
	for {
		t.Log("========")
		time.Sleep(1 * time.Second)
		return
	}
}
