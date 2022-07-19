/**
 * @Author: EDZ
 * @Description:
 * @File: iconnmanager
 * @Version: 1.0.0
 * @Date: 2022/7/14 9:29
 */

package ziface

type IConnManager interface {
	//新增链接
	AddConnection(connection IConnection) error
	//获取链接
	GetConnection(id uint32) (IConnection, error)
	//删除链接
	DelConnection(id uint32) error
	//清空链接
	ClearConnection() error
}
