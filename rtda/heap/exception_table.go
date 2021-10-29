package heap

import "gvm/classfile"

type ExceptionTable []*ExceptionHandler

type ExceptionHandler struct {
	startPc   int       // try{} 语句块的第一条指令
	endPc     int       // try{} 语句块的下一条指令
	handlerPc int       // 异常处理代码位置
	catchType *ClassRef // 如果为nil（class文件中是0），表示可以处理所以异常，用来实现finally子句
}

// 把class文件中的异常处理表转换成 ExceptionTable 类型
func newExceptionTable(entries []*classfile.ExceptionTableEntry, cp *ConstantPool) ExceptionTable {
	table := make([]*ExceptionHandler, len(entries))
	for i, entry := range entries {
		table[i] = &ExceptionHandler{
			startPc:   int(entry.StartPc()),
			endPc:     int(entry.EndPc()),
			handlerPc: int(entry.HandlerPc()),
			catchType: getCatchType(uint(entry.CatchType()), cp),
		}
	}

	return table
}

// 从运行时常量池中查找类符号引用
func getCatchType(index uint, cp *ConstantPool) *ClassRef {
	if index == 0 {
		return nil // catch all
	}
	return cp.GetConstant(index).(*ClassRef)
}

// 如果位于startPc和endPc之间的指令抛出异常x，且x为X（或X的子类）的实例，就返回对应的异常处理器
func (self ExceptionTable) findExceptionHandler(exClass *Class, pc int) *ExceptionHandler {
	for _, handler := range self {
		// jvms: The start_pc is inclusive and end_pc is exclusive
		if pc >= handler.startPc && pc < handler.endPc {
			if handler.catchType == nil {
				return handler
			}
			catchClass := handler.catchType.ResolvedClass()
			if catchClass == exClass || catchClass.IsSuperClassOf(exClass) {
				return handler
			}
		}
	}
	return nil
}
