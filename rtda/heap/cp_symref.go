package heap

// 类符号引用
type SymRef struct {
	cp        *ConstantPool // 此引用属于的类的常量池
	className string        // 此引用对应的类名
	class     *Class
}

func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}
	return self.class
}

// 如果 类d中的引用n 要加载c类型 则需要获取d的类加载器加载c 否则会抛出异常
func (self *SymRef) resolveClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.class = c
}
