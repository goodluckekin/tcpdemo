/**
 * @Author: 玩家模块
 * @Description:
 * @File: player
 * @Version: 1.0.0
 * @Date: 2022/7/20 10:31
 */

package player

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"math/rand"
	"zinx/ziface"
)

type Player struct {
	Pid  int64 //玩家id
	Conn ziface.IConnection
	X    float32 //x坐标
	Y    float32 //y坐标
	Z    float32 //2d不存在
	V    float32 //角度
}

func NewPlayer(conn ziface.IConnection) *Player {
	return &Player{
		Pid:  rand.Int63(),
		Conn: conn,
		X:    float32(50 + rand.Intn(10)), //初始化出生地，随机数
		Y:    float32(40 + rand.Intn(10)),
		Z:    0,
		V:    0,
	}
}

//给客户端发送消息
func (p *Player) SendMsg(msgId uint32, message proto.Message) error {
	buf, err := proto.Marshal(message)
	if err != nil {
		fmt.Println("protoc marshal error:", err)
		return err
	}

	if err := p.Conn.Send(msgId, buf); err != nil {
		return err
	}
	return nil
}
