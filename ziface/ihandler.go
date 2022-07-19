/**
 * @Author: EDZ
 * @Description:
 * @File: ihandler
 * @Version: 1.0.0
 * @Date: 2022/7/8 15:37
 */

package ziface

type IHandler interface {

	//路由处理方法
	MsgHandler(req IRequest) error

	//增加路由
	AddRouter(id uint32, r IRouter) error
}
