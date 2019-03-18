package rpcServer

import "reflect"

// 反射出来的方法对象
type StructMethodObject struct {
	// 反射出来的对应方法对象
	Method reflect.Value

	// 反射出来的方法的输入参数的类型集合
	InTypes []reflect.Type

	// 反射出来的方法的输出参数的类型集合
	OutTypes []reflect.Type
}

// 创建反射对象
// 参数：
// method：反射出来的方法
// intypes：方法参数
// outTypes：方法返回值
// 返回值：
// 1.方法对象
func NewStructMethodObject(method reflect.Value, intypes, outTypes []reflect.Type) *StructMethodObject {
	return &StructMethodObject{
		Method:   method,
		InTypes:  intypes,
		OutTypes: outTypes,
	}
}
