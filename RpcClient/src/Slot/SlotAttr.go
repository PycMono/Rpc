package Slot

import (
	"fmt"
	"moqikaka.com/Rpc/RpcClient/src/Enum"
)

// 属性信息
type SlotAttr struct {
	// 属性字典
	fightDict map[int]int
}

// 创建新的卡槽属性信息
func NewSlotAttr() *SlotAttr {
	newObj := &SlotAttr{}
	newObj.fightDict = newObj.initFightDict()

	return newObj
}

// 重置属性
func (this *SlotAttr) Reset()  {
	this.fightDict = this.initFightDict()
}

// 累加战斗属性
func (this *SlotAttr)AddValue(attrDict map[int]int)  {
	for key,value := range attrDict{
		if _,ok:=this.fightDict[key];!ok{
				fmt.Println(fmt.Sprintf("this.fightDict不存key=%d的枚举",key) )
		}

		this.fightDict[key] += value
	}
}

// 累加战斗属性
func (this * SlotAttr)AddValue2(key,value int)  {
	if _,ok:=this.fightDict[key];!ok{
		fmt.Println(fmt.Sprintf("this.fightDict不存key=%d的枚举",key) )
	}

	this.fightDict[key] += value
}

// 初始化属性字典
func (this *SlotAttr)initFightDict()map[int]int  {
	result := make(map[int]int)
	result[Enum.ATK] = 0
	result[Enum.HP] = 0
	result[Enum.MaxHP] = 0

	return result
}