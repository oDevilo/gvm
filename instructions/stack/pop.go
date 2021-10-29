package stack

import (
	"gvm/instructions/base"
	"gvm/rtda"
)

// Pop the top operand stack value
type POP struct{ base.NoOperandsInstruction }

func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// pop指令只能用于弹出int、float等占用一个操作数栈位置的变量。double和long变量在操作数栈中占据两个位置，需要使用pop2指令弹出
type POP2 struct{ base.NoOperandsInstruction }

func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
