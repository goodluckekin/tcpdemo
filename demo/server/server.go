/**
 * @Author: ekin
 * @Description:启动一个tcp服务器
 * @File: server
 * @Version: 1.0.0
 * @Date: 2022/7/5 14:06
 */

package main

import (
	"fmt"
	"zinx/demo/core/bus"
	"zinx/demo/core/player"
	"zinx/demo/pb/msg"
	"zinx/demo/router"
	"zinx/ziface"
	"zinx/znet"
)

func main() {
	srv := znet.NewServer("test")
	r := router.NewRouter()
	srv.AddRouter(1, r)
	srv.SetStartConnectHook(func(connection ziface.IConnection) {
		fmt.Println("connection start hook")

		//新建玩家
		p := player.NewPlayer(connection)
		fmt.Println("====>new player id:", p.Pid, " is coming <=====")

		//链接关联玩家id
		connection.SetProperty("playerId", p.Pid)

		//加入到格子中
		bus.Wm.AddPlayer(p)

		//测试发送消息
		m := &msg.SyncIDMsg{
			Pid: p.Pid,
		}
		if err := p.SendMsg(10, m); err != nil {
			fmt.Println("player pid:", p.Pid, "sending msg error:", err)
		}

		//测试发送广播消息
		bus.Wm.Brocast(20, m)
	})
	srv.SetStopConnectHook(func(connection ziface.IConnection) {
		fmt.Println("connection stop hook")

		//移除玩家
		pid := connection.GetProperty("playerId").(int64)
		p := bus.Wm.GetPlayer(int(pid))
		bus.Wm.RemovePlayer(p)
	})
	srv.Serve()
}
