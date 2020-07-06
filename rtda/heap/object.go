package heap

type Object struct {
	class *Class
	data  interface{} // 表示存放普通对象或者数组，实际存放的是Slots变量
}

// 创建对象
func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}

// Get方法
func (self *Object) Class() *Class {
	return self.class
}

func (self *Object) Fields() Slots {
	return self.data.(Slots) //针对普通对象
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}

// reflection
func (self *Object) GetRefVar(name, descriptor string) *Object {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	return slots.GetRef(field.slotId)
}
func (self *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	slots.SetRef(field.slotId, ref)
}
