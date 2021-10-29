package native

import "gvm/rtda"

type NativeMethod func(frame *rtda.Frame)

// 本地方法注册表 k: 方法的唯一标识 v: 具体方法的本地实现
var registry = map[string]NativeMethod{}

func emptyNativeMethod(frame *rtda.Frame) {
	// do nothing
}

// 类名、方法名、方法描述符一起确定一个唯一方法
func Register(className, methodName, methodDescriptor string, method NativeMethod) {
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = method
}

// 查找本地方法实现，如果找不到则返回nil
func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := className + "~" + methodName + "~" + methodDescriptor
	if method, ok := registry[key]; ok {
		return method
	}
	// Object/Thread/System等类通过registerNatives()的本地方法来注册其他本地方法
	// registerNatives用于初始化注册，所以后面还需调用的时候直接返回空
	if methodDescriptor == "()V" {
		// initIDs 为了 FileInputStream的initIDs方法
		if methodName == "registerNatives" || methodName == "initIDs" {
			return emptyNativeMethod
		}
	}
	return nil
}
