package references

import (
	"gvm/instructions/base"
	"gvm/rtda"
	"gvm/rtda/heap"
)

// Invoke instance method; dispatch based on class
// 与 invokeinterface 区别在于 此指令调用方的this引用执行某个类（或其子类）的实例，可以利用vtable
// 而 invokeinterface 可以是实现了该接口的类实例，无法使用vtable技术，会慢一点
type INVOKE_VIRTUAL struct{ base.Index16Instruction }

func (self *INVOKE_VIRTUAL) Execute(frame *rtda.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	if resolvedMethod.IsProtected() &&
		resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&
		!ref.Class().IsSubClassOf(currentClass) {

		if !(ref.Class().IsArray() && resolvedMethod.Name() == "clone") {
			panic("java.lang.IllegalAccessError")
		}
	}

	// 从对象的类中查找真正要调用的方法，如果找不到或者是抽象方法则抛出异常
	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}
