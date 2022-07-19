/**
 * @Author: EDZ
 * @Description:
 * @File: iconnection
 * @Version: 1.0.0
 * @Date: 2022/7/5 14:35
 */

package ziface

import "net"

type IConnection interface {
	//启动
	Start()

	//结束
	Stop()

	//发送消息
	Send(id uint32, data []byte) error

	//获取链接id
	GetConnID() uint32

	//获取链接
	GetConn() *net.TCPConn
}
