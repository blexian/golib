package reflect

import (
	"testing"
)

// 示例函数
func ExampleFunc() {}

type FuncTestStruct struct{}

func (m FuncTestStruct) ExampleMethod()         {}
func (m *FuncTestStruct) PointerExampleMethod() {}

func TestGetFuncName(t *testing.T) {
	// 测试匿名函数
	anonymous := func() {}
	// 测试nil函数
	var nilFunc func()
	fts := FuncTestStruct{}
	cases := map[string]any{
		"github.com/blexian/golib/pkgs/reflect.ExampleFunc":                               ExampleFunc,
		"github.com/blexian/golib/pkgs/reflect.FuncTestStruct.ExampleMethod-fm":           fts.ExampleMethod,
		"github.com/blexian/golib/pkgs/reflect.(*FuncTestStruct).PointerExampleMethod-fm": fts.PointerExampleMethod,
		"github.com/blexian/golib/pkgs/reflect.TestGetFuncName.func1":                     anonymous,
		"<nil func>": nilFunc,
	}
	index := 0
	for except, funcCase := range cases {
		index++
		result := GetFuncName(funcCase)
		if result != except {
			t.Errorf("serial number %d, except %v, but %v", index, except, result)
		}
	}
}
