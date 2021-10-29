package loads

import (
	"gvm/instructions/base"
	"gvm/rtda"
)

// Load int from local variable
type ILOAD struct{ base.Index8Instruction }

func (self *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, self.Index)
}

// 第0个变量压入操作数栈
type ILOAD_0 struct{ base.NoOperandsInstruction }

func (self *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

// 第1个变量压入操作数栈
type ILOAD_1 struct{ base.NoOperandsInstruction }

func (self *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

type ILOAD_2 struct{ base.NoOperandsInstruction }

func (self *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

type ILOAD_3 struct{ base.NoOperandsInstruction }

func (self *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}

func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}
