/**
 * @Author: EDZ
 * @Description:
 * @File: irouter
 * @Version: 1.0.0
 * @Date: 2022/7/5 15:19
 */

package ziface

type IRouter interface {
	//处理签的钩子方法
	PreHandler(request IRequest) error

	//处理函数
	Handler(request IRequest) error

	//处理后的函数
	PostHandler(request IRequest) error

	//错误处理
	ErrorHandler(err error, request IRequest)
}
