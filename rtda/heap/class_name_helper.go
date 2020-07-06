package heap

/*
 类型描述符
 */
var primitiveTypes = map[string]string{
	"void":    "V",
	"boolean": "Z",
	"byte":    "B",
	"short":   "S",
	"int":     "I",
	"long":    "J",
	"char":    "C",
	"float":   "F",
	"double":  "D",
}

/*
 [XXX -> [[XXX
 int -> [I
 XXX -> [LXXX;
 */
func getArrayClassName(className string) string {
	return "[" + toDescriptor(className)
}

/*
 [[XXX -> [XXX
 [LXXX; -> XXX
 [I -> int
 */
func getComponentClassName(className string) string {
	if className[0] == '[' {
		componentTypeDescriptor := className[1:]
		return toClassName(componentTypeDescriptor)
	}
	panic("Not array: " + className)
}

/*
 [XXX -> [XXX
 int  -> I
 XXX  -> LXXX;
 */
func toDescriptor(className string) string {
	//判断是否为数组类型
	if className[0] == '[' {
		return className //直接返回数组类型描述符
	}
	//判断是否为基本类型
	if d, ok := primitiveTypes[className]; ok {
		return d //返回基本类型描述符
	}
	//返回引用类型描述符
	return "L" + className + ";"
}

/*
 [XXX  -> [XXX
 LXXX; -> XXX
 I     -> int
 */
func toClassName(descriptor string) string {
	if descriptor[0] == '[' {
		return descriptor
	}
	if descriptor[0] == 'L' {
		return descriptor[1:len(descriptor) - 1]
	}
	for className, d := range primitiveTypes{
		if d == descriptor {
			return className
		}
	}
	panic("Invalid descriptor: " + descriptor)
}
