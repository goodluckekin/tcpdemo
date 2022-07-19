/**
 * @Author: ekin
 * @Description:封包与拆包
 * @File: idatapack
 * @Version: 1.0.0
 * @Date: 2022/7/6 9:40
 */

package ziface

type IDataPack interface {
	//获取协议数据包head长度
	GetDataPackHeadLen() uint32
	//封包
	Pack(msg IMessage) ([]byte, error)
	//拆包
	Unpack([]byte) (IMessage, error)
}
