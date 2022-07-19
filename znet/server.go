/**
 * @Author: EDZ
 * @Description:
 * @File: server
 * @Version: 1.0.0
 * @Date: 2022/7/5 13:48
 */

package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

var _ ziface.IServer = (*Server)(nil)

type Server struct {
	name             string
	ipVersion        string
	ip               string
	port             uint32
	handlers         ziface.IHandler
	maxPoolSize      int //携程池最大数量
	maxTaskLen       int //最大任务数
	taskPool         []chan ziface.IRequest
	connManager      *ConnManager
	startConnectHook ziface.HookFunc //链接创建hook
	stopConnectHook  ziface.HookFunc //链接关闭hook
}

func NewServer(name string) *Server {
	return &Server{
		name:        name,
		ipVersion:   "tcp4",
		ip:          "127.0.0.1",
		port:        8999,
		handlers:    NewHandler(),
		maxPoolSize: 10,
		maxTaskLen:  10,
		taskPool:    nil,
		connManager: NewConnManager(),
	}
}

//创建携程池
func (s *Server) StartTaskPool() {
	s.taskPool = make([]chan ziface.IRequest, s.maxPoolSize)
	for i := 0; i < s.maxPoolSize; i++ {
		s.taskPool[i] = make(chan ziface.IRequest, s.maxTaskLen)
		go func(i int) {
			s.StartWorker(i)
		}(i)
	}
}

//创建一个worker
func (s *Server) StartWorker(i int) {
	fmt.Printf("worker %d is starting!\n", i)
	for req := range s.taskPool[i] {

		//调用路由处理消息
		err := s.handlers.MsgHandler(req)
		if err != nil {
			conn := req.GetConnection()
			if err := conn.Send(req.GetMsgId(), []byte(err.Error())); err != nil {
				fmt.Println("send msg error", err)
				return
			}
		}
	}
}

//任务投递 根据链接id取余，投递到不同的worker
func (s *Server) Dispatch(req ziface.IRequest) error {
	conn := req.GetConnection()
	num := int(conn.GetConnID()) % s.maxPoolSize
	fmt.Println("request dispatch to worker ", num)
	if ch := s.taskPool[num]; ch == nil {
		return fmt.Errorf("request dispatch error:%d", conn.GetConnID())
	} else {
		s.taskPool[num] <- req
	}
	return nil
}

//启动tcp连接
func (s *Server) Start() {
	fmt.Printf("【server】 %s:%d start ...\n", s.ip, s.port)
	addr, err := net.ResolveTCPAddr(s.ipVersion, fmt.Sprintf("%s:%d", s.ip, s.port))
	if err != nil {
		fmt.Printf("【server】 ip resolve error %v \n", err)
		return
	}

	listener, err := net.ListenTCP(s.ipVersion, addr)
	defer listener.Close()
	if err != nil {
		fmt.Printf("【server】 listen error %v \n", err)
		return
	}

	var i int
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Printf("【server】 acceptTcp error %v \n", err)
			continue
		}

		i++
		notifyConn := NewConnection(s, uint32(i), conn, s.handlers)
		notifyConn.Start()
	}

}

//设置创建链接钩子函数
func (s *Server) SetStartConnectHook(f ziface.HookFunc) {
	s.startConnectHook = f
}

func (s *Server) GetStartConnectHook() ziface.HookFunc {
	return s.startConnectHook
}

//设置结束链接钩子函数
func (s *Server) SetStopConnectHook(f ziface.HookFunc) {
	s.stopConnectHook = f
}

func (s *Server) GetStopConnectHook() ziface.HookFunc {
	return s.stopConnectHook
}

//获取链接管理
func (s *Server) GetConnManager() ziface.IConnManager {
	return s.connManager
}

//关闭资源
func (s *Server) Stop() {

	//关闭携程池
	for _, e := range s.taskPool {
		close(e)
	}
}

//启动服务
func (s *Server) Serve() {
	s.StartTaskPool() //创建工作协成池
	go s.Start()
	defer s.Stop()
	select {}
}

//添加路由
func (s *Server) AddRouter(id uint32, r ziface.IRouter) {
	err := s.handlers.AddRouter(id, r)
	if err != nil {
		panic(err)
	}
}
