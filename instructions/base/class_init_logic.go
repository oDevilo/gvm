package base

import (
	"gvm/rtda"
	"gvm/rtda/heap"
)

/**
类初始化在下列情况下触发
1. 执行new指令创建类实例，但类还没被初始化
2. 执行putstatic、getstatic指令存取类的静态变量，但声明该字段的类没被初始化
3. 执行invokestatic调用类的静态方法，但声明该方法的类还没初始化
4. 当初始化一个类，如果类的超类没被初始化，先初始化超类
5. 执行某些反射操作时
*/
func InitClass(thread *rtda.Thread, class *heap.Class) {
	class.StartInit()
	scheduleClinit(thread, class)
	initSuperClass(thread, class)
}

func scheduleClinit(thread *rtda.Thread, class *heap.Class) {
	clinit := class.GetClinitMethod()
	if clinit != nil {
		// exec <clinit>
		newFrame := thread.NewFrame(clinit)
		thread.PushFrame(newFrame)
	}
}

func initSuperClass(thread *rtda.Thread, class *heap.Class) {
	if !class.IsInterface() {
		superClass := class.SuperClass()
		if superClass != nil && !superClass.InitStarted() {
			InitClass(thread, superClass)
		}
	}
}
