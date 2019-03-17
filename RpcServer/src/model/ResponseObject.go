package model

const (
	con_Default_MethodName = "SendMessage"
)

// Socket服务器的响应对象
type ResponseObject struct {
	// 响应结果的状态值
	*ResultStatus

	// 响应结果的数据
	Data interface{} `json:"Data,omitempty"`

	// 响应结果对应的请求的方法名称
	MethodName string
}

func (this *ResponseObject) SetResultStatus(rs *ResultStatus) *ResponseObject {
	this.ResultStatus = rs

	return this
}

func (this *ResponseObject) SetData(data interface{}) *ResponseObject {
	this.Data = data

	return this
}

func (this *ResponseObject) SetMethodName(methodName string) *ResponseObject {
	this.MethodName = methodName

	return this
}

func NewResponseObject() *ResponseObject {
	return &ResponseObject{
		ResultStatus: Success,
		Data:         nil,
		MethodName:   "",
	}
}
