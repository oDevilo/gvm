package base

import (
	"gvm/rtda"
	"gvm/rtda/heap"
)

/**
执行方法
会创建一个新的帧并推入虚拟机栈顶，然后传递参数
*/
func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {
	// 创建新的帧并推入虚拟机栈
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)

	// 参数数量不一定是java代码中参数数量，如long和double类型占用两个位置
	// 如果是实例方法，最前面会有个隐藏参数this，静态方法没有this
	// 将需要的参数从操作栈中弹出放入被调用方法的局部变量表
	argSlotCount := int(method.ArgSlotCount())
	if argSlotCount > 0 {
		for i := argSlotCount - 1; i >= 0; i-- {
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}
}
