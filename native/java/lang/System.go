package lang

import (
	"gvm/instructions/base"
	"gvm/native"
	"gvm/rtda"
	"gvm/rtda/heap"
	"runtime"
	"time"
)

const jlSystem = "java/lang/System"

// System类初始化分为两个阶段：
// 1. 由类初始化方法完成，在此方法中 registerNatives 会注册其他本地方法
// 2. 由vm完成，vm会调用 System.initializeSystemClass() （由sun.misc.VM的initialize方法调用）
/**
System中的 in out err 都是在此位置创建
FileInputStream fdIn = new FileInputStream(FileDescriptor.in);
FileOutputStream fdOut = new FileOutputStream(FileDescriptor.out);
FileOutputStream fdErr = new FileOutputStream(FileDescriptor.err);
setIn0(new BufferedInputStream(fdIn));
setOut0(newPrintStream(fdOut, props.getProperty("sun.stdout.encoding")));
setErr0(newPrintStream(fdErr, props.getProperty("sun.stderr.encoding")));
*/
func init() {
	native.Register(jlSystem, "arraycopy", "(Ljava/lang/Object;ILjava/lang/Object;II)V", arraycopy)
	native.Register(jlSystem, "initProperties", "(Ljava/util/Properties;)Ljava/util/Properties;", initProperties)
	native.Register(jlSystem, "setIn0", "(Ljava/io/InputStream;)V", setIn0)
	native.Register(jlSystem, "setOut0", "(Ljava/io/PrintStream;)V", setOut0)
	native.Register(jlSystem, "setErr0", "(Ljava/io/PrintStream;)V", setErr0)
	native.Register(jlSystem, "currentTimeMillis", "()J", currentTimeMillis)
}

// public static native void arraycopy(Object src, int srcPos, Object dest, int destPos, int length)
// (Ljava/lang/Object;ILjava/lang/Object;II)V
func arraycopy(frame *rtda.Frame) {
	// 从局部变量获取到5个入参
	vars := frame.LocalVars()
	src := vars.GetRef(0)
	srcPos := vars.GetInt(1)
	dest := vars.GetRef(2)
	destPos := vars.GetInt(3)
	length := vars.GetInt(4)
	// 源数组和目标数组都不能为null
	if src == nil || dest == nil {
		panic("java.lang.NullPointerException")
	}
	// 源数组和目标数组必须兼容
	if !checkArrayCopy(src, dest) {
		panic("java.lang.ArrayStoreException")
	}
	// 检查数组越界等问题
	if srcPos < 0 || destPos < 0 || length < 0 ||
		srcPos+length > src.ArrayLength() ||
		destPos+length > dest.ArrayLength() {
		panic("java.lang.IndexOutOfBoundsException")
	}

	heap.ArrayCopy(src, dest, srcPos, destPos, length)
}

// 两个数组兼容性检测
func checkArrayCopy(src, dest *heap.Object) bool {
	srcClass := src.Class()
	destClass := dest.Class()
	// 必须都是数组
	if !srcClass.IsArray() || !destClass.IsArray() {
		return false
	}
	// 数据类型
	if srcClass.ComponentClass().IsPrimitive() || destClass.ComponentClass().IsPrimitive() {
		return srcClass == destClass
	}
	return true
}

// private static native Properties initProperties(Properties props);
// (Ljava/util/Properties;)Ljava/util/Properties;
func initProperties(frame *rtda.Frame) {
	vars := frame.LocalVars()
	props := vars.GetRef(0)

	stack := frame.OperandStack()
	stack.PushRef(props)

	// public synchronized Object setProperty(String key, String value)
	setPropMethod := props.Class().GetInstanceMethod("setProperty", "(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	thread := frame.Thread()
	for key, val := range _sysProps() {
		jKey := heap.JString(frame.Method().Class().Loader(), key)
		jVal := heap.JString(frame.Method().Class().Loader(), val)
		ops := rtda.NewOperandStack(3)
		ops.PushRef(props)
		ops.PushRef(jKey)
		ops.PushRef(jVal)
		shimFrame := rtda.NewShimFrame(thread, ops)
		thread.PushFrame(shimFrame)

		base.InvokeMethod(shimFrame, setPropMethod)
	}
}

func _sysProps() map[string]string {
	return map[string]string{
		"java.version":         "1.8.0",
		"java.vendor":          "jvm.go",
		"java.vendor.url":      "https://github.com/zxh0/jvm.go",
		"java.home":            "todo",
		"java.class.version":   "52.0",
		"java.class.path":      "todo",
		"java.awt.graphicsenv": "sun.awt.CGraphicsEnvironment",
		"os.name":              runtime.GOOS,   // todo
		"os.arch":              runtime.GOARCH, // todo
		"os.version":           "",             // todo
		"file.separator":       "/",            // todo os.PathSeparator
		"path.separator":       ":",            // todo os.PathListSeparator
		"line.separator":       "\n",           // todo
		"user.name":            "",             // todo
		"user.home":            "",             // todo
		"user.dir":             ".",            // todo
		"user.country":         "CN",           // todo
		"file.encoding":        "UTF-8",
		"sun.stdout.encoding":  "UTF-8",
		"sun.stderr.encoding":  "UTF-8",
	}
}

// private static native void setIn0(InputStream in);
// (Ljava/io/InputStream;)V
func setIn0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	in := vars.GetRef(0)

	sysClass := frame.Method().Class()
	sysClass.SetRefVar("in", "Ljava/io/InputStream;", in)
}

// 设置输出流
// private static native void setOut0(PrintStream out);
// (Ljava/io/PrintStream;)V
func setOut0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	out := vars.GetRef(0)

	sysClass := frame.Method().Class()
	sysClass.SetRefVar("out", "Ljava/io/PrintStream;", out)
}

// private static native void setErr0(PrintStream err);
// (Ljava/io/PrintStream;)V
func setErr0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	err := vars.GetRef(0)

	sysClass := frame.Method().Class()
	sysClass.SetRefVar("err", "Ljava/io/PrintStream;", err)
}

// public static native long currentTimeMillis();
// ()J
func currentTimeMillis(frame *rtda.Frame) {
	millis := time.Now().UnixNano() / int64(time.Millisecond)
	stack := frame.OperandStack()
	stack.PushLong(millis)
}
