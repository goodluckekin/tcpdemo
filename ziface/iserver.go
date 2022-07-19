/**
 * @Author: EDZ
 * @Description:
 * @File: iserver
 * @Version: 1.0.0
 * @Date: 2022/7/5 13:48
 */

package ziface

type HookFunc func(connection IConnection)

type IServer interface {
	//启动
	Start()

	//停止
	Stop()

	//运行
	Serve()

	//新增路由
	AddRouter(uint32, IRouter)

	//启动工作池
	StartTaskPool()

	//启动一个worker
	StartWorker(i int)

	//投递任务到协程池
	Dispatch(req IRequest) error

	//增加一些生命周期hook
	SetStartConnectHook(f HookFunc)
	GetStartConnectHook() HookFunc
	SetStopConnectHook(f HookFunc)
	GetStopConnectHook() HookFunc

	//获取连接管理
	GetConnManager() IConnManager
}
