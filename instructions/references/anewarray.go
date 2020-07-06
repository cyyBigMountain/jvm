package references

import (
	"jvm/instructions/base"
	"jvm/rtda"
	"jvm/rtda/heap"
)

// 引用数组结构体
type ANEW_ARRAY struct{
	base.Index16Instruction //包含运行时常量池类符号引用索引
}

func (self *ANEW_ARRAY) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	componentClass := classRef.ResolvedClass()

	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	arrClass := componentClass.ArrayClass()
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}
