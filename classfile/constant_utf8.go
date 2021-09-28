package classfile

/**
字符串在class文件中是以MUTF-8（Modified UTF-8）方式编码而不是标准的UTF-8编码
MUTF-8编码方式和UTF-8大致相同，但并不兼容
差别有两点：
一是null字符（代码点U+0000）会被编码成2字节：0xC0、0x80；
二是补充字符（Supplementary Characters，代码点大于U+FFFF的Unicode字符）是按UTF-16拆分为代理对（Surrogate Pair）分别编码的
*/
// 存放MUTF-8编码的字符串
type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

/**
因为Go语言字符串使用UTF-8编码，所以如果字符串中不包含null字符或补充字符，
下面这个简化版的readMUTF8（）也是可以工作的
*/
func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
