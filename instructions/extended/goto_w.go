package extended

import (
	"gvm/instructions/base"
	"gvm/rtda"
)

// 和goto的唯一区别就是索引从2字节变成4字节

// Branch always (wide index)
type GOTO_W struct {
	offset int
}

func (self *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}
func (self *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.offset)
}
