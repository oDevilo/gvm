package references

import (
	"gvm/instructions/base"
	"gvm/rtda"
	"gvm/rtda/heap"
	"reflect"
)

// Throw exception or error
type ATHROW struct{ base.NoOperandsInstruction }

func (self *ATHROW) Execute(frame *rtda.Frame) {
	// 弹出一个异常对象引用
	ex := frame.OperandStack().PopRef()
	if ex == nil {
		panic("java.lang.NullPointerException")
	}

	thread := frame.Thread()
	// 查看是否可以找到并跳转到异常处理代码
	if !findAndGotoExceptionHandler(thread, ex) {
		// 如果找不到异常处理逻辑
		handleUncaughtException(thread, ex)
	}
}

func findAndGotoExceptionHandler(thread *rtda.Thread, ex *heap.Object) bool {
	for {
		// java中异常时上抛的，所以从当前帧开始，遍历之前的帧直到找到异常处理项
		frame := thread.CurrentFrame()
		pc := frame.NextPC() - 1

		handlerPC := frame.Method().FindExceptionHandler(ex.Class(), pc)
		if handlerPC > 0 {
			stack := frame.OperandStack()
			stack.Clear()
			stack.PushRef(ex)
			frame.SetNextPC(handlerPC)
			return true
		}

		thread.PopFrame()
		if thread.IsStackEmpty() {
			break
		}
	}
	return false
}

// todo
func handleUncaughtException(thread *rtda.Thread, ex *heap.Object) {
	// 清空java虚拟机栈，打印异常信息
	thread.ClearStack()

	jMsg := ex.GetRefVar("detailMessage", "Ljava/lang/String;")
	goMsg := heap.GoString(jMsg)
	println(ex.Class().JavaName() + ": " + goMsg)

	stes := reflect.ValueOf(ex.Extra())
	for i := 0; i < stes.Len(); i++ {
		ste := stes.Index(i).Interface().(interface {
			String() string
		})
		println("\tat " + ste.String())
	}
}
