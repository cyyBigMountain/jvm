package native

import "jvm/rtda"

/*
 定义本地方法类型
 */
type NativeMethod func(frame *rtda.Frame)

/*
 存放本地方法表
 */
var registry = map[string]NativeMethod{}

/*
 将类名，方法名，方法描述符组合，作为本地方法表的键
 */
func Register(className, methodName, methodDescriptor string, method NativeMethod)  {
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = method
}

/*
 根据类名，方法名，方法描述符组合查找本地方法表中的方法
 */
func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := className + "~" + methodName + "~" + methodDescriptor
	if method, ok := registry[key]; ok {
		return method
	}

	// 如果方法描述符为无返回，且方法名为"registerNatives"，则返回空的本地方法
	// java.lang.Object等类通过registerNatives()的本地方法来注册其他本地方法
	if methodDescriptor == "()V" && methodName == "registerNatives" {
		return emptyNativeMethod
	}

	return nil
}

func emptyNativeMethod(frame *rtda.Frame)  {
	// 空实现
}