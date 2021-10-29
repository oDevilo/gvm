package base

import "gvm/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader) // 从字节码中提取操作数
	Execute(frame *rtda.Frame)            // 执行指令逻辑
}

type NoOperandsInstruction struct{}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

type BranchInstruction struct {
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8()) // 读取一个 uint8 整数 转成unit后赋值给 index
}

type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16()) // 读取一个 uint16 整数 转成unit后赋值给 index
}
