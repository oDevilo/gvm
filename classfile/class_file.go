package classfile

import "fmt"

type ClassFile struct {
	//magic uint32 魔数
	minorVersion uint16          // 小版本号
	majorVersion uint16          // 大版本号
	constantPool ConstantPool    // 常量池
	accessFlags  uint16          // 类访问标识 接口还是类 public还是private
	thisClass    uint16          // 类名索引 -》常量池
	superClass   uint16          // 超类名索引 在Class为Object的时候为0
	interfaces   []uint16        // 接口索引表
	fields       []*MemberInfo   // 字段表
	methods      []*MemberInfo   // 方法表
	attributes   []AttributeInfo // 属性表
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr) // 解析二进制 生成classfile
	return
}

func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}
func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}
func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}
func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return ""
}

func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

/**
魔数
很多文件格式都会规定满足该格式的文件必须以某几个固定
字节开头，这几个字节主要起标识作用，叫作魔数（magic number）。
例如PDF文件以4字节“%PDF”（0x25、0x50、0x44、0x46）开头，ZIP
文件以2字节“PK”（0x50、0x4B）开头。class文件的是“0xCAFEBABE”
*/
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

/**
特定的Java虚拟机实现只能支持版本号在某个范围内的class文
件。Oracle的实现是完全向后兼容的，比如Java SE 8支持版本号为
45.0~52.0的class文件。如果版本号不在支持的范围内，Java虚拟机
实现就抛出java.lang.UnsupportedClassVersionError异常。
我们参考Java 8，支持版本号为45.0~52.0的class文件。
JDK 1.0.2 45.0 - 45.3
JDK 1.1   45.0 - 45.65525
JDK 1.2   46.0
JDK 1.3   47.0
JDK 1.4   48.0
JDK 1.5   49.0
JDK 1.6   50.0
JDK 1.7   51.0
JDK 1.8   52.0
*/
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion { // major 表示大版本号
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 { // minor 小版本号
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

// 版本号后面是常量池

// 常量池后面是类访问标志

// 类访问标志之后是两个u2类型的常量池索引 分别给出类名和超类名

// 类和超类索引后面是接口索引表 表中存放的也是常量池索引，给出该类实现的所有接口的名字

func (self *ClassFile) SourceFileAttribute() *SourceFileAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *SourceFileAttribute:
			return attrInfo.(*SourceFileAttribute)
		}
	}
	return nil
}
