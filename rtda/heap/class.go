package heap

import (
	"gvm/classfile"
	"strings"
)

type Class struct {
	accessFlags       uint16 // 对应 access_flags
	name              string // thisClassName 完全限定名 java/lang/Object的形式
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	sourceFile        string
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint    // 实例变量所占大小
	staticSlotCount   uint    // 类变量所占大小
	staticVars        Slots   // 静态变量
	initStarted       bool    // 是否初始化
	jClass            *Object // 通过此字段，Class结构体实例与一个类对象关联
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	class.sourceFile = getSourceFile(cf)
	return class
}

// 从class文件中获取源文件名
func getSourceFile(cf *classfile.ClassFile) string {
	if sfAttr := cf.SourceFileAttribute(); sfAttr != nil {
		return sfAttr.FileName()
	}
	return "Unknown" // todo
}

func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}
func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}
func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}
func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags&ACC_ANNOTATION
}
func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}

// getters
func (self *Class) AccessFlags() uint16 {
	return self.accessFlags
}
func (self *Class) Name() string {
	return self.name
}
func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}
func (self *Class) Fields() []*Field {
	return self.fields
}
func (self *Object) Data() interface{} {
	return self.data
}
func (self *Class) Methods() []*Method {
	return self.methods
}
func (self *Class) SourceFile() string {
	return self.sourceFile
}
func (self *Class) Loader() *ClassLoader {
	return self.loader
}
func (self *Class) SuperClass() *Class {
	return self.superClass
}
func (self *Class) Interfaces() []*Class {
	return self.interfaces
}
func (self *Class) StaticVars() Slots {
	return self.staticVars
}
func (self *Class) InitStarted() bool {
	return self.initStarted
}
func (self *Class) JClass() *Object {
	return self.jClass
}

func (self *Class) StartInit() {
	self.initStarted = true
}

// jvms 5.4.4
func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsPublic() ||
		self.getPackageName() == other.getPackageName()
}

func (self *Class) getPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (self *Class) GetClinitMethod() *Method {
	return self.getStaticMethod("<clinit>", "()V")
}

func (self *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range self.methods {
		if method.IsStatic() &&
			method.name == name &&
			method.descriptor == descriptor {

			return method
		}
	}
	return nil
}

func (self *Class) getMethod(name, descriptor string, isStatic bool) *Method {
	for c := self; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.IsStatic() == isStatic &&
				method.name == name &&
				method.descriptor == descriptor {

				return method
			}
		}
	}
	return nil
}

func (self *Class) getField(name, descriptor string, isStatic bool) *Field {
	for c := self; c != nil; c = c.superClass {
		for _, field := range c.fields {
			if field.IsStatic() == isStatic &&
				field.name == name &&
				field.descriptor == descriptor {

				return field
			}
		}
	}
	return nil
}

func (self *Class) GetConstructor(descriptor string) *Method {
	return self.GetInstanceMethod("<init>", descriptor)
}

func (self *Class) GetConstructors(publicOnly bool) []*Method {
	constructors := make([]*Method, 0, len(self.methods))
	for _, method := range self.methods {
		if method.isConstructor() {
			if !publicOnly || method.IsPublic() {
				constructors = append(constructors, method)
			}
		}
	}
	return constructors
}

func (self *Class) GetFields(publicOnly bool) []*Field {
	if publicOnly {
		publicFields := make([]*Field, 0, len(self.fields))
		for _, field := range self.fields {
			if field.IsPublic() {
				publicFields = append(publicFields, field)
			}
		}
		return publicFields
	} else {
		return self.fields
	}
}

func (self *Class) GetMethods(publicOnly bool) []*Method {
	methods := make([]*Method, 0, len(self.methods))
	for _, method := range self.methods {
		if !method.isClinit() && !method.isConstructor() {
			if !publicOnly || method.IsPublic() {
				methods = append(methods, method)
			}
		}
	}
	return methods
}

func (self *Class) isJlObject() bool {
	return self.name == "java/lang/Object"
}
func (self *Class) isJlCloneable() bool {
	return self.name == "java/lang/Cloneable"
}
func (self *Class) isJioSerializable() bool {
	return self.name == "java/io/Serializable"
}

func (self *Class) NewObject() *Object {
	return newObject(self)
}

func (self *Class) ArrayClass() *Class {
	arrayClassName := getArrayClassName(self.name)
	return self.loader.LoadClass(arrayClassName)
}

// java/lang/Object -> java.lang.Object
func (self *Class) JavaName() string {
	return strings.Replace(self.name, "/", ".", -1)
}

// 是否基础类型
func (self *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[self.name]
	return ok
}

func (self *Class) GetInstanceMethod(name, descriptor string) *Method {
	return self.getMethod(name, descriptor, false)
}

func (self *Class) GetStaticMethod(name, descriptor string) *Method {
	return self.getMethod(name, descriptor, true)
}

func (self *Class) GetRefVar(fieldName, fieldDescriptor string) *Object {
	field := self.getField(fieldName, fieldDescriptor, true)
	return self.staticVars.GetRef(field.slotId)
}
func (self *Class) SetRefVar(fieldName, fieldDescriptor string, ref *Object) {
	field := self.getField(fieldName, fieldDescriptor, true)
	self.staticVars.SetRef(field.slotId, ref)
}

/**
判定此Class对象所表示的类或接口与指定的Class参数所表示的类或接口是否相同,或是否是其超类或超接口
数组可以强制转换成Object类型（因为数组的超类是Object）
数组可以强制转换成Cloneable和Serializable类型（因为数组实现了这两个接口）
如果下面两个条件之一成立，类型为[]SC的数组可以强制转换成类型为[]TC的数组：
    1.TC和SC是同一个基本类型
    2.TC和SC都是引用类型，且SC可以强制转换成TC
*/
func (self *Class) IsAssignableFrom(other *Class) bool {
	s, t := other, self

	if s == t {
		return true
	}

	if !s.IsArray() {
		if !s.IsInterface() {
			// s is class
			if !t.IsInterface() {
				// t is not interface
				return s.IsSubClassOf(t)
			} else {
				// t is interface
				return s.IsImplements(t)
			}
		} else {
			// s is interface
			if !t.IsInterface() {
				// t is not interface
				return t.isJlObject()
			} else {
				// t is interface
				return t.isSuperInterfaceOf(s)
			}
		}
	} else {
		// s is array
		if !t.IsArray() {
			if !t.IsInterface() {
				// t is class
				return t.isJlObject()
			} else {
				// t is interface
				return t.isJlCloneable() || t.isJioSerializable()
			}
		} else {
			// t is array
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			return sc == tc || tc.IsAssignableFrom(sc)
		}
	}

	return false
}

func (self *Class) GetPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

// 是否某个类的子类
func (self *Class) isSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

// 是否实现某个接口
func (self *Class) isImplements(iface *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

// 是否继承某个接口
func (self *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}

// 是否某个类的父类
func (self *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(self)
}

// 是否某个类的子类
func (self *Class) IsSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

// 是否实现某个接口
func (self *Class) IsImplements(iface *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

// 是否某个接口的父接口
func (self *Class) isSuperInterfaceOf(iface *Class) bool {
	return iface.isSubInterfaceOf(self)
}
