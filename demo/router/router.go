/**
 * @Author: EDZ
 * @Description:
 * @File: router
 * @Version: 1.0.0
 * @Date: 2022/7/5 15:48
 */

package router

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

var _ ziface.IRouter = (*Router)(nil)

type Router struct {
	znet.BaseRouter
}

func NewRouter() *Router {
	return &Router{}
}

func (r Router) Handler(req ziface.IRequest) error {
	fmt.Printf("【router】 handler...\n")
	err := req.GetConnection().Send(req.GetMsgId(), req.GetData())
	return err
}

func (r Router) ErrorHandler(err error, req ziface.IRequest) {
	fmt.Printf("【router】 error handler...err:%v \n", err)
}
