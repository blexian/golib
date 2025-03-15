package reflect

import (
	"fmt"
	"reflect"
	"strings"
)

// GetVarType 通过反射获取变量的类型
func GetVarType(v any) reflect.Type {
	return reflect.TypeOf(v)
}

// GetVarTypeName 通过反射获取变量不包含包名的类型名称
// nil 对象 返回 nil
// 支持多层指针类型对象，例如
// **int对象 返回 **int
// **(type TestStruct struct) 返回 **TestStruct
func GetVarTypeName(v any) string {
	t := reflect.TypeOf(v)
	if t == nil {
		return "nil"
	}

	// 递归解析指针层级
	ptrLevel := 0
	for t.Kind() == reflect.Ptr {
		ptrLevel++
		t = t.Elem()
	}

	// 构建类型名称
	var builder strings.Builder
	builder.WriteString(strings.Repeat("*", ptrLevel))
	builder.WriteString(getBaseTypeName(t))

	return builder.String()
}

// getAbsoluteBaseTypeName 获取基础包含包名类型名称（处理结构体等自定义类型）
// 直接处理指针无法获取低层类型信息
func getAbsoluteBaseTypeName(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Struct:
		return fmt.Sprintf("%s.%s", t.PkgPath(), t.Name())
	default:
		return t.String()
	}
}

// getBaseTypeName 获取基础类型名称（处理结构体等自定义类型）
// 直接处理指针无法获取低层类型信息
func getBaseTypeName(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Struct:
		return fmt.Sprintf("%s", t.Name())
	default:
		return t.String()
	}
}

// GetVarTypeKind 通过反射获取变量的类型kind，kind表示大类如struct,interface
func GetVarTypeKind(v any) reflect.Kind {
	return reflect.ValueOf(v).Kind()
}

// VarIsPointer 判断对象是否为指针
func VarIsPointer(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.Ptr
}

func test() {
	v := reflect.ValueOf(nil)
	v.Elem()
	fmt.Println(v.Kind().String())
}
