package rpcServer

import (
	"fmt"
	"reflect"
	"testing"
)

type Student struct {
	age  int32
	name string
}

func (this *Student) GetName() {

}

func (this *Student) GetAge() int32 {
	return this.age
}

func (this *Student) setName() int32 {
	return this.age
}

func (this *Student) SetAge(age int32) {
	this.age = age
}

func TestReflect(t *testing.T) {
	structObject := new(Student)
	reflectType := reflect.TypeOf(structObject)
	reflectValue := reflect.ValueOf(structObject)

	//reflectTypeStr := reflectType.String()
	//reflectTypeArr := strings.Split(reflectTypeStr, ".")
	//fmt.Println(reflectTypeStr)
	//fmt.Println(reflectTypeArr[len(reflectTypeArr)-1])

	// 获取方法名字
	for i := 0; i < reflectType.NumMethod(); i++ {
		methodName := reflectType.Method(i).Name
		method := reflectValue.MethodByName(methodName)
		methodType := method.Type()
		for i := 0; i < methodType.NumIn(); i++ {
			fmt.Println(methodType.In(i))
		}

		fmt.Println("----------------------")
	}

}
