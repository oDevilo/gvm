package classfile

type MarkerAttribute struct{}

// 指出类、接口、字段、方法已不建议使用
type DeprecatedAttribute struct{ MarkerAttribute }

// 标记源文件中不存在、由编译器生成的类成员，引入Synthetic属性主要是为了支持嵌套类和嵌套接口
type SyntheticAttribute struct{ MarkerAttribute }

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
