/**
 * @Author: EDZ
 * @Description:
 * @File: handler
 * @Version: 1.0.0
 * @Date: 2022/7/8 15:37
 */

package znet

import (
	"errors"
	"zinx/ziface"
)

var _ ziface.IHandler = (*Handler)(nil)

type Handler struct {
	routers map[uint32]ziface.IRouter
}

func NewHandler() *Handler {
	return &Handler{
		routers: make(map[uint32]ziface.IRouter),
	}
}

func (h *Handler) MsgHandler(req ziface.IRequest) error {
	msgId := req.GetMsgId()
	if r, exist := h.routers[msgId]; !exist {
		return errors.New("router not found")
	} else {
		if err := r.PreHandler(req); err != nil {
			return err
		}
		if err := r.Handler(req); err != nil {
			return err
		}
		if err := r.PostHandler(req); err != nil {
			return err
		}
	}
	return nil
}

func (h *Handler) AddRouter(id uint32, r ziface.IRouter) error {
	if _, exist := h.routers[id]; exist {
		return errors.New("router exists")
	}

	h.routers[id] = r
	return nil
}
