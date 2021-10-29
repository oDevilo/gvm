package misc

import (
	"gvm/instructions/base"
	"gvm/native"
	"gvm/rtda"
)

func init() {
	native.Register("sun/misc/VM", "initialize", "()V", initialize)
}

// 调用System类的方法
// private static native void initialize();
// ()V
func initialize(frame *rtda.Frame) {
	classLoader := frame.Method().Class().Loader()
	jlSysClass := classLoader.LoadClass("java/lang/System")
	initSysClass := jlSysClass.GetStaticMethod("initializeSystemClass", "()V")
	base.InvokeMethod(frame, initSysClass)
}
