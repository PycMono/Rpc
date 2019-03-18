package model

// 客服端请求对象
type RequestObject struct {
	// 以下属性是由客户端直接传入的，可以直接反序列化直接得到的
	// 请求的模块名称
	ModuleName string

	// 请求的方法名称
	MethodName string

	// 请求的参数数组
	Parameters []interface{}
}