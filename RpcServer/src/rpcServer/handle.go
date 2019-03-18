package rpcServer

import (
	"encoding/json"
	"fmt"
	"moqikaka.com/Rpc/RpcServer/src/model"
	"moqikaka.com/goutil/debugUtil"
	"moqikaka.com/goutil/logUtil"
	"net"
	"time"
)

var (
	PACKAGE_NAME = "rpcServer"
)

// 处理连接
func handleConn(conn net.Conn) {
	methodName := "handleConn"

	// 创建新的客服端
	clientObj := newClient(conn)
	Register(clientObj) // 注册到客服端管理对象中

	// 打印客服端数量
	fmt.Println(fmt.Sprintf("handleConn客服端连接数量%d", GetClientCount()))

	// 是否退出和客服端发送消息
	ch := make(chan bool, 1)
	// 处理消息发送
	go handleSendData(clientObj, ch)

	defer func() {
		logUtil.ErrorLog(fmt.Sprintf("删除玩家clientID=%d", clientObj.GetID()))
		clientObj.Quit()
		URegister(clientObj) // 删除玩家连接
		ch <- true
	}()

	// 无限循环，接收消息，处理消息
	for {
		readBytes := make([]byte, 1024)
		// 读取数据
		n, err := clientObj.conn.Read(readBytes)
		if err != nil {
			logUtil.ErrorLog(fmt.Sprintf("%s_%s,read data error,err=%s", PACKAGE_NAME, methodName, err))
			break
		}

		// 将数据追加到读取数据末尾,为什么要用:n,可能读取到的数据没有1024那么大
		clientObj.appendReceiveData(readBytes[:n])

		// 数据处理
		handleReceiveData(clientObj)
	}
}

// 处理数据读取
// 参数：
// clientObj：客服端对象
// 返回值：无
func handleReceiveData(clientObj *Client) {
	// 循环读取数据
	for {
		receiveData, exists := clientObj.getReceiveData()
		if !exists {
			break
		}

		if len(receiveData) == 0 {
			// 收到心跳包

			continue
		}

		requestObj := new(model.RequestObject)
		err := json.Unmarshal(receiveData, requestObj)
		if err != nil {
			fmt.Println(fmt.Sprintf("处理数据报错msg=%v", err))
			continue
		}

		fmt.Println(fmt.Sprintf("cliendID=%d,接收到了消息%s", clientObj.GetID(), requestObj.MethodName))

		// 数据处理
		CallFunction(requestObj)
	}
}

// 处理消息发送
func handleSendData(clientObj *Client, ch chan bool) {
	for {
		select {
		case <-ch: // 收到ch的命令，表示clientObj已经断开
			debugUtil.Println("------------handleSendData收到handleConn发来的消息，goroutine退出------------")
			return
		default:
			if responseObject, exists := clientObj.getSendData(); exists {
				debugUtil.Println("------------exists data------------")
				if err := clientObj.sendMessage(responseObject); err != nil {
					debugUtil.Println("------------send data failed，goroutine退出------------")
					return
				}

				debugUtil.Println("------------send data successfully------------")
			} else {
				time.Sleep(5 * time.Minute)

				// 找不到数据则判断是否需要新增一条空数据
				//if getCurrMillsecond()-unixMillsecond > intervalMillsecond {
				//	clientObj.appendSendData(NewResponseObject())
				//}
			}
		}
	}
}

type Person struct {
	Name string
	Age  int
}
