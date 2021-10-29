package vm

import (
	"fmt"
	"gvm/classpath"
	"gvm/instructions/base"
	"gvm/rtda"
	"gvm/rtda/heap"
	"strings"
)

type JVM struct {
	cmd         *Cmd
	classLoader *heap.ClassLoader
	mainThread  *rtda.Thread
}

func NewJVM(cmd *Cmd) *JVM {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)
	return &JVM{
		cmd:         cmd,
		classLoader: classLoader,
		mainThread:  rtda.NewThread(),
	}
}

func (self *JVM) Start() {
	self.initVM()
	self.execMain()
}

// 虚拟机先加载 sun.misc.VM 类 执行其初始化方法
func (self *JVM) initVM() {
	vmClass := self.classLoader.LoadClass("sun/misc/VM")
	base.InitClass(self.mainThread, vmClass)
	interpret(self.mainThread, self.cmd.verboseInstFlag)
}

// 加载主类，执行其main方法
func (self *JVM) execMain() {
	className := strings.Replace(self.cmd.Class, ".", "/", -1)
	mainClass := self.classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod == nil {
		fmt.Printf("Main method not found in class %s\n", self.cmd.Class)
		return
	}

	argsArr := self.createArgsArray()
	frame := self.mainThread.NewFrame(mainMethod)
	frame.LocalVars().SetRef(0, argsArr)
	self.mainThread.PushFrame(frame)
	interpret(self.mainThread, self.cmd.verboseInstFlag)
}

func (self *JVM) createArgsArray() *heap.Object {
	stringClass := self.classLoader.LoadClass("java/lang/String")
	argsLen := uint(len(self.cmd.args))
	argsArr := stringClass.ArrayClass().NewArray(argsLen)
	jArgs := argsArr.Refs()
	for i, arg := range self.cmd.args {
		jArgs[i] = heap.JString(self.classLoader, arg)
	}
	return argsArr
}
