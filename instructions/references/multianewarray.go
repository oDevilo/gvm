package references

import (
	"gvm/instructions/base"
	"gvm/rtda"
	"gvm/rtda/heap"
)

// 多维数组

// Create new multidimensional array
type MULTI_ANEW_ARRAY struct {
	index      uint16
	dimensions uint8
}

func (self *MULTI_ANEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	self.index = reader.ReadUint16()
	self.dimensions = reader.ReadUint8()
}
func (self *MULTI_ANEW_ARRAY) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(uint(self.index)).(*heap.ClassRef)
	arrClass := classRef.ResolvedClass()

	stack := frame.OperandStack()
	counts := popAndCheckCounts(stack, int(self.dimensions))
	arr := newMultiDimensionalArray(counts, arrClass)
	stack.PushRef(arr)
}

//  获取每个维度的长度
func popAndCheckCounts(stack *rtda.OperandStack, dimensions int) []int32 {
	counts := make([]int32, dimensions)
	for i := dimensions - 1; i >= 0; i-- {
		counts[i] = stack.PopInt()
		if counts[i] < 0 {
			panic("java.lang.NegativeArraySizeException")
		}
	}

	return counts
}

// 创建多维数组
func newMultiDimensionalArray(counts []int32, arrClass *heap.Class) *heap.Object {
	// 创建第一维度数组
	count := uint(counts[0])
	arr := arrClass.NewArray(count)

	if len(counts) > 1 {
		refs := arr.Refs()
		// 为此维度每个元素赋值数组 递归操作
		for i := range refs {
			refs[i] = newMultiDimensionalArray(counts[1:], arrClass.ComponentClass())
		}
	}

	return arr
}
