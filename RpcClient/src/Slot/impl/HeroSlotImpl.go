package impl

import (
	"fmt"
	"moqikaka.com/RpcClient/src/Enum"
	"moqikaka.com/RpcClient/src/Slot"
)

func init()  {
	Slot.RegisterFunction( &HeroSlotImpl{})
}

type  HeroSlotImpl struct {

}

func (this *HeroSlotImpl)GetSlotEnum()int32  {
	return  Enum.Hero
}

// 获取枚举
func (this * HeroSlotImpl)Println(){
	fmt.Println("计算武将属性：Hero")
}