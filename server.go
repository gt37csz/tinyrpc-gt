// Copyright 2022 <mzh.scnu@qq.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tinyrpc

import (
	"net"
	"net/rpc"

	"github.com/zehuamama/tinyrpc/codec"
	"github.com/zehuamama/tinyrpc/serializer"
)

// 基本上是对net/rpc包的包装，添加了序列化字段
// Server rpc server based on net/rpc implementation
type Server struct {
	*rpc.Server
	serializer.Serializer
}

// NewServer Create a new rpc server
func NewServer(opts ...Option) *Server {
	options := options{
		serializer: serializer.Proto,
	}
	for _, option := range opts {
		option(&options)
	}

	return &Server{&rpc.Server{}, options.serializer}
}

// Register register rpc function
func (s *Server) Register(rcvr interface{}) error {
	return s.Server.Register(rcvr)
}

// RegisterName register the rpc function with the specified name
func (s *Server) RegisterName(name string, rcvr interface{}) error {
	return s.Server.RegisterName(name, rcvr)
}

// 每次有客户端连上来，服务器端都会开启一个goroutine，该goroutine就会只负责这一个连接
// Serve start service
func (s *Server) Serve(lis net.Listener) {
	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}
		// 它会在for循环中重复的执行server.readRequest(codec)
		// 经过debug，该方法退出循环的一种方式是，它没有阻塞，当conn中没有请求数据可以读时
		// 反返回一个EOF的error，然后就break退出循环了。但是在退出该方法，关闭连接之前
		// 会阻塞直到之前所有call都执行完毕并发送给客户端。
		go s.Server.ServeCodec(codec.NewServerCodec(conn, s.Serializer))
	}
}
