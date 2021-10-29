package lang

import (
	"fmt"
	"gvm/native"
	"gvm/rtda"
	"gvm/rtda/heap"
)

const jlThrowable = "java/lang/Throwable"

// 记录java虚拟机栈帧信息
type StackTraceElement struct {
	fileName   string // 类所在文件名
	className  string // 声明方法的类名
	methodName string // 方法名
	lineNumber int    // 帧正在执行哪行代码
}

func (self *StackTraceElement) String() string {
	return fmt.Sprintf("%s.%s(%s:%d)",
		self.className, self.methodName, self.fileName, self.lineNumber)
}

func init() {
	native.Register(jlThrowable, "fillInStackTrace", "(I)Ljava/lang/Throwable;", fillInStackTrace)
}

// private native Throwable fillInStackTrace(int dummy);
// (I)Ljava/lang/Throwable;
func fillInStackTrace(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	frame.OperandStack().PushRef(this)

	stes := createStackTraceElements(this, frame.Thread())
	this.SetExtra(stes)
}

func createStackTraceElements(tObj *heap.Object, thread *rtda.Thread) []*StackTraceElement {
	// 由于栈顶两帧正在执行 fillInStackTrace(int) fillInStackTrace() 所以需要跳过
	// 下面几帧正在执行异常类的构造函数，所以也要跳过，具体跳过多少得看异常类的继承层次
	skip := distanceToObject(tObj.Class()) + 2
	frames := thread.GetFrames()[skip:]
	stes := make([]*StackTraceElement, len(frames))
	for i, frame := range frames {
		stes[i] = createStackTraceElement(frame)
	}
	return stes
}

// 计算需要异常类继承层数
func distanceToObject(class *heap.Class) int {
	distance := 0
	for c := class.SuperClass(); c != nil; c = c.SuperClass() {
		distance++
	}
	return distance
}

func createStackTraceElement(frame *rtda.Frame) *StackTraceElement {
	method := frame.Method()
	class := method.Class()
	return &StackTraceElement{
		fileName:   class.SourceFile(),
		className:  class.JavaName(),
		methodName: method.Name(),
		lineNumber: method.GetLineNumber(frame.NextPC() - 1),
	}
}
