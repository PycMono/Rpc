package Hero

import "moqikaka.com/RpcClient/src/Slot"

// 武将对象
type  Hero struct {
	// 武将ID
	HeroID string

	// 战斗属性
	*Slot.SlotAttr

	// SlotEnum:枚举模块对应的具体模块的属性，用来查看数据
	SlotAttrDict map[int32]*Slot.SlotAttr
}

func NewHero() *Hero {
	newObj:=&Hero{
		HeroID:"ss",
		SlotAttr:Slot.NewSlotAttr(),
		SlotAttrDict: map[int32]*Slot.SlotAttr{},
	}

	return newObj
}

//
func initSlotAttrDict()map[int32]*Slot.SlotAttr  {
	result:=make(map[int32]*Slot.SlotAttr)

	return result
}
