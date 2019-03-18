package main

import (
	"moqikaka.com/Rpc/RpcServer/src/rpcServer"
	"sync"
)

var (
	wg sync.WaitGroup
)

// 初始化函数
func init() {
	// 设置WaitGroup需要等待的数量，只要有一个服务器出现错误都停止服务器
	wg.Add(1)
}

func main() {
	go rpcServer.Start(&wg)

	wg.Wait()
}
