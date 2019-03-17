package RpcClient

import (
	"fmt"
	"sync"
)

var (
	// 客服端连接集合
	ClientList []*Client

	// 锁对象
	reMutex sync.RWMutex
)

func init() {
	ClientList = make([]*Client,0)
}

// 注册客服端对象
// 参数：
// id：客服端唯一ID
// client：客服端对象
func Register(client *Client) {
	reMutex.RLock()
	defer reMutex.RUnlock()

	ClientList=append(ClientList,client)

	// 打印客服端数量
	fmt.Println(fmt.Sprintf("Register连接数量%d",len(ClientList)))
}
