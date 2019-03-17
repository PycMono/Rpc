package Slot

import "fmt"

var(
	MethodDict map[int32]ISlot
)

func init()  {
	MethodDict=make(map[int32]ISlot)
}

func RegisterFunction(structObject ISlot) {
	if _,ok:=MethodDict[structObject.GetSlotEnum()];ok{
		fmt.Println(fmt.Sprintf("RegisterFunction函数中Enum=%d的枚举已经存在",structObject.GetSlotEnum()))
		return
	}

	MethodDict[structObject.GetSlotEnum()] = structObject
}

func CallFunction()  {
		for _,value:=range MethodDict {
			value.Println()
	}
}