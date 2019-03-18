package rpcServer

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"moqikaka.com/Rpc/RpcServer/src/model"
	"moqikaka.com/goutil/fileUtil"
	"moqikaka.com/goutil/intAndBytesUtil"
	"moqikaka.com/goutil/logUtil"
	"moqikaka.com/goutil/timeUtil"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// 全局客户端的id，从1开始进行自增
	globalClientId int32 = 0

	// 字节的大小端顺序
	byterOrder = binary.LittleEndian
)

const (
	// 包头长度
	con_HEADER_LENTH = 4
)

// 客服端对象
type Client struct {
	// ID
	id int32

	// 连接对象
	conn net.Conn

	// 玩家ID
	playerID string

	// 上次活跃时间
	activeTime time.Time

	// 锁对象
	mutex sync.Mutex

	// 发送的数据
	receiveData []byte

	// 接收的数据
	sendData []*model.ResponseObject
}

// 获取客服端ID
func (this *Client) GetID() int32 {
	return this.id
}

// 记录日志
// 参数：
// msg：消息信息
// 返回值：
// 无
func (this *Client) WriteLog(msg string) {
	fileUtil.WriteFile("LOG", this.getRemoteAddr(), true,
		timeUtil.Format(time.Now(), "yyyy-MM-dd HH:mm:ss"),
		" ",
		fmt.Sprintf("client:%s", this.String()),
		" ",
		msg,
		"\r\n",
		"\r\n",
	)
}

// 格式化数据
// 参数：无
// 返回值：
// 格式化的数据
func (this *Client) String() string {
	return fmt.Sprintf("{Id:%d, RemoteAddr:%s, activeTime:%s, playerId:%s}", this.id, this.getRemoteAddr(), timeUtil.Format(this.activeTime, "yyyy-MM-dd HH:mm:ss"), this.playerID)
}

// 关闭连接方法
func (this *Client) Quit() {
	this.conn.Close()
}

// 追加接收的数据
// 参数：
// receiveData：接收的数据
func (this *Client) appendReceiveData(receiveData []byte) {
	this.receiveData = append(this.receiveData, receiveData...)
	this.activeTime = time.Now()
}

// 读取数据
// 参数：
// 1.读取值
// 2.是否存在数据
// 返回值：
// 读取的数据
func (this *Client) getReceiveData() ([]byte, bool) {
	// 判断是否包含头部信息
	if len(this.receiveData) < con_HEADER_LENTH {
		return nil, false
	}

	// 获取头部信息
	header := this.receiveData[:con_HEADER_LENTH]

	// 将头部数据转换为内部的长度
	contentLength := intAndBytesUtil.BytesToInt32(header, byterOrder)

	// 判断长度是否满足
	if len(this.receiveData) < con_HEADER_LENTH+int(contentLength) {
		return nil, false
	}

	// 提取消息内容
	content := this.receiveData[con_HEADER_LENTH : con_HEADER_LENTH+contentLength]

	// 将对应的数据截断，以得到新的数据
	this.receiveData = this.receiveData[con_HEADER_LENTH+contentLength:]

	return content, true
}

// 追加发送的数据
// sendDataItemObj:待发送数据项
// 返回值：无
func (this *Client) appendSendData(responseObj *model.ResponseObject) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	this.sendData = append(this.sendData, responseObj)
}

// 获取待发送的数据
// 返回值：
// 待发送数据项
// 是否含有有效数据
func (this *Client) getSendData() (responseObj *model.ResponseObject, exists bool) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	// 如果没有数据则直接返回
	if len(this.sendData) == 0 {
		return
	}

	// 取出第一条数据,并为返回值赋值
	responseObj = this.sendData[0]
	exists = true

	// 删除已经取出的数据
	this.sendData = this.sendData[1:]

	return
}

// 获取远程连接地址和端口
// 参数：无
// 返回值
// 客服端地址和端口
func (this *Client) getRemoteAddr() string {
	return this.conn.RemoteAddr().String()
}

// 发送字节数组消息
// responseObj:返回值对象
func (this *Client) sendMessage(responseObj *model.ResponseObject) error {
	beforeTime := time.Now().Unix()

	// 序列化发送的数据
	content, err := json.Marshal(responseObj)
	if err != nil {
		logUtil.NormalLog("序列化response数据失败", logUtil.Error)
		return errors.New("序列化response数据失败")
	}

	contentStr := string(content)
	this.WriteLog(contentStr)

	// 进行zlib压缩
	//baseConfig := config.GetBaseConfig()
	//if baseConfig.IfCompressClientData {
	//	content, err = zlibUtil.Compress(content, zlib.DefaultCompression)
	//	if err != nil {
	//		logUtil.NormalLog(fmt.Sprintf("压缩Data出错，错误信息为：%s", err), logUtil.Error)
	//		return errors.New("压缩Data出错")
	//	}
	//}

	// 获得数据内容的长度
	contentLength := len(content)

	// 将长度转化为字节数组
	header := intAndBytesUtil.Int32ToBytes(int32(contentLength), byterOrder)

	// 将头部与内容组合在一起
	message := append(header, content...)

	// 发送消息
	if _, err = this.conn.Write(message); err != nil {
		logUtil.NormalLog(fmt.Sprintf("发送消息,%s,出现错误：%s", contentStr, err), logUtil.Error)
		return err
	}

	// 如果发送的时间超过3秒，则记录下来
	if time.Now().Unix()-beforeTime > 3 {
		logUtil.NormalLog(fmt.Sprintf("消息Size:%d, UseTime:%d", contentLength, time.Now().Unix()-beforeTime), logUtil.Warn)
	}

	return err
}

// 新建客户端对象
// conn：连接对象
// 返回值：客户端对象的指针
func newClient(conn net.Conn) *Client {
	// 获得自增的id值
	getIncrementId := func() int32 {
		atomic.AddInt32(&globalClientId, 1)
		return globalClientId
	}

	return &Client{
		id:          getIncrementId(),
		conn:        conn,
		receiveData: make([]byte, 0, 1024),
		activeTime:  time.Now(),
		playerID:    "",
	}
}
