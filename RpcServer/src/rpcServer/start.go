package rpcServer

import (
	"errors"
	"fmt"
	"net"
	"sync"
)

// 启动服务器
// 参数：
// wg：WaitGroup
func Start(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	// 监听地址
	listener, err := net.Listen("tcp", "10.254.0.162:8989")
	if err != nil {
		panic(errors.New(fmt.Sprintf("Listen Error: %s", err)))
	}

	msg := fmt.Sprintf("listener success，addr=%s", listener.Addr())
	fmt.Print(msg)

	for {
		// 阻塞，等待新的请求过来
		conn, err := listener.Accept()
		if err != nil {
			fmt.Print(fmt.Sprintf("Accept error,err=%s", err))
			continue
		}

		// 处理连接
		go handleConn(conn)
	}
}
