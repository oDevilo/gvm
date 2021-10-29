package heap

import "gvm/classfile"

// 针对引用是类的情况
type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodRefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}

func (self *MethodRef) ResolvedMethod() *Method {
	if self.method == nil {
		self.resolveMethodRef()
	}
	return self.method
}

// jvms8 5.4.3.3
func (self *MethodRef) resolveMethodRef() {
	// 如果类d想访问类c的某个方法，需要先解析符号引用获取类c
	d := self.cp.class
	c := self.ResolvedClass()
	// 如果c是接口
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 找到c中对应的方法
	method := lookupMethod(c, self.name, self.descriptor)
	// 如果找不到方法
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	// 如果类d没有访问权限
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.method = method
}

func lookupMethod(class *Class, name, descriptor string) *Method {
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}
