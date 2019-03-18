package rpcServer

import (
	"fmt"
	"moqikaka.com/Rpc/RpcServer/src/model"
	"moqikaka.com/goutil/debugUtil"
	"moqikaka.com/goutil/logUtil"
	"reflect"
	"strings"
)

const (
	// 包名字
	con_MODULE_NAME = "rpcServer.reflect"

	// 供客户端访问的模块的后缀
	con_ModuleSuffix = "Module"

	// 定义用于分隔模块名称和方法名称的分隔符
	con_DelimeterOfObjAndMethod = "_"
)

var (
	// 定义存放所有方法映射的变量
	methodDict = make(map[string]*StructMethodObject)
)

// 获取结构体对象名字
// 参数：
// structType：结构体类型
// 返回值：
// 1.结构体名
func getStructName(structType reflect.Type) string {
	structStr := structType.String()
	structArray := strings.Split(structStr, ",")

	return structArray[len(structArray)-1]
}

// 获取参数和返回值类型集合
// 参数：
// method：方法类型
// 返回值：
// 1.参数类型
// 2.返回值
func getMethodInAndOutParams(method reflect.Value) (inTypes []reflect.Type, outTypes []reflect.Type) {
	methodType := method.Type()
	for i := 0; i < methodType.NumIn(); i++ {
		inTypes = append(inTypes, methodType.In(i))
	}

	for i := 0; i < methodType.NumOut(); i++ {
		outTypes = append(outTypes, methodType.Out(i))
	}

	return
}

// 获取全类名
// 参数：
// structName：结构体名字
// methodName：方法名字
// 返回值：
// 1.拼接后的方法名字
func getFullMethodName(structName, methodName string) string {
	return fmt.Sprintf("%s%s%s", structName, con_DelimeterOfObjAndMethod, methodName)
}

// 注册方法
// 参数：
// structObject：结构体对象
func RegisterFunction(structObject interface{}) {
	structType := reflect.TypeOf(structObject)
	structValue := reflect.ValueOf(structObject)
	if structType.Kind() != reflect.Struct {
		logUtil.ErrorLog(fmt.Sprintf("%s.%s反射出来的类型不是结构体类型type=%s", con_MODULE_NAME, "RegisterFunction", structType.Kind()))
		return
	}

	// 获取结构体名
	structName := getStructName(structType)
	for i := 0; i < structType.NumMethod(); i++ {
		// 获取方法名字
		methodName := structType.Method(i).Name

		method := structValue.MethodByName(methodName)
		// 获取方法参数
		inTypes, outTypes := getMethodInAndOutParams(method)

		// todo pyc判断方法参数和返回值是否满足条件

		// 添加到列表中
		methodDict[getFullMethodName(structName, methodName)] = NewStructMethodObject(method, inTypes, outTypes)

		debugUtil.Println(fmt.Sprintf("%s_%s注册成功,当前共%d个方法", structName, methodName, len(methodDict)))
	}
}

// 调用方法
// 参数：
// requestObject：请求对象
// 返回值：
// 1.返回值对象
func CallFunction(requestObject *model.RequestObject) *model.ResponseObject {

	return nil
}
