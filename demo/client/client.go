/**
 * @Author: 客户端
 * @Description:
 * @File: client
 * @Version: 1.0.0
 * @Date: 2022/7/5 14:07
 */

package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
	"zinx/demo/pb/player"
	"zinx/znet"
)

func main() {
	fmt.Printf("【client】 start ...\n")
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	defer conn.Close()
	if err != nil {
		fmt.Printf("【client】 connect error:%v \n", err)
		return
	}

	//数据封包类
	dp := znet.NewDataPack()
	wg := sync.WaitGroup{}
	wg.Add(3)

	//读取消息
	go func() {
		defer wg.Done()
		for {
			//读取头部信息
			headBuf := make([]byte, dp.GetDataPackHeadLen())
			if _, err := io.ReadFull(conn, headBuf); err != nil {
				fmt.Printf("【client】 recv data error:%v \n", err)
				return
			}

			if msg, err := dp.Unpack(headBuf); err != nil {
				fmt.Printf("【client】 head Unpack error:%v \n", err)
				return
			} else {

				//读写消息
				if msg.GetMsgLen() > 0 {
					dataBuf := make([]byte, msg.GetMsgLen())
					_, err = io.ReadFull(conn, dataBuf)
					if err != nil {
						fmt.Printf("【client】 recv data error:%v \n", err)
						return
					}

					msg.SetMsgData(dataBuf)

					//反序列化获取内容
					var p player.Info
					if err := proto.Unmarshal(dataBuf, &p); err != nil {
						fmt.Println("proto unmarshal error", err)
					}

					fmt.Printf("【client】 recv msgId:%d data:%s length:%d info:%+v\n", msg.GetMsgId(), string(msg.GetMsgData()), msg.GetMsgLen(), p)
				}
			}
		}
	}()

	//模拟写消息
	go func() {
		defer wg.Done()
		for {

			//测试发送pb数据包
			info := &player.Info{
				Id:   1,
				Name: "测试",
				Phone: []player.Phone{
					player.Phone_PC,
				},
			}

			infoBytes, err := proto.Marshal(info)
			if err != nil {
				fmt.Println("【client】 marshal error", err)
				continue
			}

			//回写测试
			msg := znet.NewMessage(1, infoBytes)
			buf, err := dp.Pack(msg)
			if err != nil {
				fmt.Printf("【client】 pack data error:%v \n", err)
				return
			}

			_, err = conn.Write(buf)
			if err != nil {
				fmt.Printf("【client】 send data error:%v \n", err)
				return
			}

			time.Sleep(1 * time.Second)
		}
	}()

	//启动pprof服务端口
	go func() {
		defer wg.Done()
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	wg.Wait()
}
