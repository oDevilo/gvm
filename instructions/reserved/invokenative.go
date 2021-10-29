package reserved

import (
	"gvm/instructions/base"
	"gvm/native"
	_ "gvm/native/java/io"
	_ "gvm/native/java/lang" // 调用init函数
	_ "gvm/native/java/security"
	_ "gvm/native/java/util/concurrent/atomic"
	_ "gvm/native/sun/io"
	_ "gvm/native/sun/misc"
	_ "gvm/native/sun/reflect"
	"gvm/rtda"
)

// Invoke native method
type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	// 根据类名、方法名、方法描述符从本地方法注册表中查找本地方法实现
	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}
