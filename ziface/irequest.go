/**
 * @Author: EDZ
 * @Description:
 * @File: irequest
 * @Version: 1.0.0
 * @Date: 2022/7/5 15:11
 */

package ziface

type IRequest interface {
	//获取链接
	GetConnection() IConnection
	//获取数据
	GetData() []byte
	//获取数据id
	GetMsgId() uint32
}
