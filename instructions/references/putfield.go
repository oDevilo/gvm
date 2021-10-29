package references

import (
	"gvm/instructions/base"
	"gvm/rtda"
	"gvm/rtda/heap"
)

// Set field in object
type PUT_FIELD struct{ base.Index16Instruction }

func (self *PUT_FIELD) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()

	// 必须是实例属性
	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 如果是final字段则只能在构造函数中初始化
	if field.IsFinal() {
		if currentClass != field.Class() || currentMethod.Name() != "<init>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		val := stack.PopInt()
		ref := self.popRef(stack)
		ref.Fields().SetInt(slotId, val)
	case 'F':
		val := stack.PopFloat()
		ref := self.popRef(stack)
		ref.Fields().SetFloat(slotId, val)
	case 'J':
		val := stack.PopLong()
		ref := self.popRef(stack)
		ref.Fields().SetLong(slotId, val)
	case 'D':
		val := stack.PopDouble()
		ref := self.popRef(stack)
		ref.Fields().SetDouble(slotId, val)
	case 'L', '[':
		val := stack.PopRef()
		ref := self.popRef(stack)
		ref.Fields().SetRef(slotId, val)
	default:
		// todo
	}
}

func (self *PUT_FIELD) popRef(stack *rtda.OperandStack) *heap.Object {
	ref := stack.PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	return ref
}
