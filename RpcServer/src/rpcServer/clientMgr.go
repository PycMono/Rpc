package rpcServer

import (
	"fmt"
	"moqikaka.com/goutil/logUtil"
	"sync"
)

var (
	// 客服端连接集合
	ClientDict map[int32]*Client

	// 锁对象
	reMutex sync.RWMutex
)

func init() {
	ClientDict = make(map[int32]*Client, 1024)
}

// 注册客服端对象
// 参数：
// id：客服端唯一ID
// client：客服端对象
func Register(client *Client) {
	reMutex.RLock()
	defer reMutex.RUnlock()

	if _, ok := ClientDict[client.GetID()]; ok {
		logUtil.ErrorLog(fmt.Sprintf("clientManager.Register：ID=%d", client.GetID()))
		return
	}

	ClientDict[client.GetID()] = client
}

// 移除客服端对象
// 参数：
// client：客服端对象
func URegister(client *Client) {
	reMutex.RLock()
	defer reMutex.RUnlock()

	if _, ok := ClientDict[client.GetID()]; !ok {
		logUtil.ErrorLog(fmt.Sprintf("clientManager.URegister：ID=%d", client.GetID()))
		return
	}

	delete(ClientDict, client.GetID())
}

// 获取客服端对象
// 参数：
// id：客服端唯一id
// 返回值：
// 1.客服端对象
// 2.是否存在
func GetClient(id int32) (*Client, bool) {
	reMutex.RLock()
	defer reMutex.RUnlock()

	if client, ok := ClientDict[id]; ok {
		return client, true
	}

	logUtil.ErrorLog(fmt.Sprintf("clientManager.GetClient：ID=%d", id))
	return nil, false
}

// 获取客服端连接数量
// 参数：无
// 返回值：
// 客服端数量
func GetClientCount() int {
	reMutex.RLock()
	defer reMutex.RUnlock()

	return len(ClientDict)
}
