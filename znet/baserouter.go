/**
 * @Author: EDZ
 * @Description:
 * @File: BaseRouter
 * @Version: 1.0.0
 * @Date: 2022/7/5 15:21
 */

package znet

import (
	"fmt"
	"zinx/ziface"
)

var _ ziface.IRouter = (*BaseRouter)(nil)

type BaseRouter struct{}

func NewBaseRouter() *BaseRouter {
	return new(BaseRouter)
}

func (b *BaseRouter) PreHandler(request ziface.IRequest) error {
	return nil
}

func (b *BaseRouter) Handler(request ziface.IRequest) error {
	return nil
}

func (b *BaseRouter) PostHandler(request ziface.IRequest) error {
	return nil
}

func (b *BaseRouter) ErrorHandler(err error, request ziface.IRequest) {
	fmt.Printf("【router】 error handler...err:%v \n", err)
}
