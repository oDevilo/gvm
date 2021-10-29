package references

import (
	"gvm/instructions/base"
	"gvm/rtda"
	"gvm/rtda/heap"
)

// 执行静态方法
type INVOKE_STATIC struct{ base.Index16Instruction }

func (self *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	// 方法必须是静态方法
	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 如果方法对应的类未初始化，先初始化类
	// 方法不能是初始化方法（此规则由类文件验证器保障）
	class := resolvedMethod.Class()
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	base.InvokeMethod(frame, resolvedMethod)
}
