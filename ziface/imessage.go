/**
 * @Author: EDZ
 * @Description:
 * @File: imessage
 * @Version: 1.0.0
 * @Date: 2022/7/6 9:28
 */

package ziface

type IMessage interface {
	//获取消息id
	GetMsgId() uint32
	//获取消息长度
	GetMsgLen() uint32
	//获取消息
	GetMsgData() []byte

	//设置消息id
	SetMsgId(id uint32)
	//设置消息长度
	SetMsgLen(len uint32)
	//设置消息
	SetMsgData(data []byte)
}
