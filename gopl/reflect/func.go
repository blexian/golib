package reflect

import (
	"reflect"
	"runtime"
)

// GetFuncName 获取函数名称（包含包路径和接收者类型）
func GetFuncName(fn interface{}) string {
	// 检查输入是否为函数类型
	v := reflect.ValueOf(fn)
	if v.Kind() != reflect.Func {
		return "<not a func>"
	}

	// 获取函数指针（PC）
	pc := v.Pointer()
	if pc == 0 {
		return "<nil func>"
	}

	// 获取函数元信息
	funcInfo := runtime.FuncForPC(pc)
	if funcInfo == nil {
		return "<unknown func>"
	}

	// 返回完整函数名称
	return funcInfo.Name()
}
