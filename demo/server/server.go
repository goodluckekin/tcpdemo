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
	})
	srv.SetStopConnectHook(func(connection ziface.IConnection) {
		fmt.Println("connection stop hook")
	})
	srv.Serve()
}
