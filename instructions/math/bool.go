package math

import (
	"gvm/instructions/base"
	"gvm/rtda"
)

// 布尔运算指令 令只能操作int和long变量，分为按位与（and）、按位或（or）、按位异或（xor）3

// Boolean AND int
type IAND struct{ base.NoOperandsInstruction }

func (self *IAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}

// Boolean AND long
type LAND struct{ base.NoOperandsInstruction }

func (self *LAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 & v2
	stack.PushLong(result)
}

// Boolean OR int
type IOR struct{ base.NoOperandsInstruction }

func (self *IOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 | v2
	stack.PushInt(result)
}

// Boolean OR long
type LOR struct{ base.NoOperandsInstruction }

func (self *LOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 | v2
	stack.PushLong(result)
}

// Boolean XOR int
type IXOR struct{ base.NoOperandsInstruction }

func (self *IXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	result := v1 ^ v2
	stack.PushInt(result)
}

// Boolean XOR long
type LXOR struct{ base.NoOperandsInstruction }

func (self *LXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	result := v1 ^ v2
	stack.PushLong(result)
}
