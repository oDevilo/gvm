package control

import (
	"gvm/instructions/base"
	"gvm/rtda"
)

// 返回void
type RETURN struct{ base.NoOperandsInstruction }

func (self *RETURN) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame()
}

// 返回引用类型
type ARETURN struct{ base.NoOperandsInstruction }

// 将从当前帧的操作栈中弹出返回值，放入上一帧中的操作栈
func (self *ARETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	ref := currentFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(ref)
}

// 返回double
type DRETURN struct{ base.NoOperandsInstruction }

func (self *DRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(val)
}

// 返回float
type FRETURN struct{ base.NoOperandsInstruction }

func (self *FRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(val)
}

// 返回int
type IRETURN struct{ base.NoOperandsInstruction }

func (self *IRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(val)
}

// 返回long类型
type LRETURN struct{ base.NoOperandsInstruction }

func (self *LRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(val)
}
