package RpcClient

import (
	"encoding/json"
	"fmt"
	"moqikaka.com/goutil/logUtil"
	"net"
	"sync"
	"time"
	"moqikaka.com/Rpc/RpcServer/src/model"
)

func StartClient(wg sync.WaitGroup)  {
	// 连接对象
	conn,err:= net.DialTimeout("tcp","192.168.1.174:8989", 2*time.Second)
	if err!=nil{
		fmt.Println(fmt.Sprintf("RpcClient.StartClient连接报错,err=%v",err))
	}

	clientObj:=NewClient(conn)
	Register(clientObj)
	handleSendData(clientObj)// 消息发送

	defer func() {
		conn.Close()
		clientObj = nil
		wg.Done()
	}()

	// 读取数据
	// 死循环，不断地读取数据，解析数据，发送数据
	for {
		// 先读取数据，每次读取1024个字节
		readBytes := make([]byte, 1024)

		// Read方法会阻塞，所以不用考虑异步的方式
		n, err := conn.Read(readBytes)
		if err != nil {
			break
		}

		//// 将读取到的数据追加到已获得的数据的末尾
		clientObj.appendContent(readBytes[:n])
		// 已经包含有效的数据，处理该数据
		handleClient(clientObj)
	}
}

func handleClient(clientObj *Client) {
	for {
		id, content, ok := clientObj.getValidMessage()
		if !ok {
			break
		}

		// 处理数据，如果长度为0则表示心跳包
		if len(content) == 0 {
			continue
		} else {
			_=id
			//handleMessage(id, content)
		}
	}
}

func handleSendData(clientObj *Client) {
	for {
		time.Sleep(time.Second * 3)

		person := &Person{
			Name: "张三",
			Age:  10,
		}

		parameters := make([]interface{}, 0)
		parameters=append(parameters,person)

		requestObj := &model.RequestObject{
			MethodName: "11",
			ModuleName: "fdsaf",
			Parameters:parameters,
		}

		if b, err := json.Marshal(requestObj); err != nil {
			logUtil.NormalLog(fmt.Sprintf("序列化请求数据%v出错", requestObj), logUtil.Error)
		} else {
			// 发送数据
			if clientObj != nil {
				clientObj.sendByteMessage(1, b)
			}
		}

	}
}

type Person struct {
	Name string
	Age int
}
