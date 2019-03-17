package impl

import (
	"fmt"
	"moqikaka.com/RpcClient/src/Enum"
	"moqikaka.com/RpcClient/src/Slot"
)

func init()  {
	Slot.RegisterFunction( NewEquipSlotImpl())
}

// 装备实现类
type  EquipSlotImpl struct {
	
}

// 获取枚举
func (this * EquipSlotImpl)GetSlotEnum() int32 {
	return Enum.Equip
}

// 获取枚举
func (this * EquipSlotImpl)Println(){
	fmt.Println("计算装备属性：Equip")
}

func NewEquipSlotImpl()*EquipSlotImpl  {
	return  &EquipSlotImpl{}
}
