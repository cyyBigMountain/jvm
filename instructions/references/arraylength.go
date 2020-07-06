package references

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

//获取数组长度结构体
type ARRAY_LENGTH struct{
	base.NoOperandsInstruction //无操作数码
}

func (self *ARRAY_LENGTH) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrRef := stack.PopRef()
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}

	arrLen := arrRef.ArrayLength()
	stack.PushInt(arrLen)
}