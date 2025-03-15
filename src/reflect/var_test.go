package reflect

import (
	"reflect"
	"testing"
)

type TestStruct struct {
	Name string
}

func TestGetVarTypeName(t *testing.T) {
	testStruct := TestStruct{}
	testStructPointer := &testStruct
	testStructPtrPtr := &testStructPointer
	cases := map[string]any{
		"nil":          nil,
		"int":          0,
		"uint":         uint(0),
		"int32":        '我',
		"TestStruct":   testStruct,
		"*TestStruct":  testStructPointer,
		"**TestStruct": testStructPtrPtr,
		"": struct {
			Name string
		}{
			Name: "darren",
		},
	}
	index := 0
	for typeName, variables := range cases {
		index++
		vt := GetVarTypeName(variables)
		if vt != typeName {
			t.Errorf("serial number %d except %s, but %s", index, typeName, vt)
		}
	}
}

func TestGetAbsoluteBaseTypeName(t *testing.T) {
	testStruct := TestStruct{}
	cases := map[string]any{
		"int":   0,
		"uint":  uint(0),
		"int32": '我',
		"github.com/blexian/golib/pkgs/reflect.TestStruct": testStruct,
	}
	index := 0
	for typeName, variables := range cases {
		index++
		vt := getAbsoluteBaseTypeName(reflect.TypeOf(variables))
		if vt != typeName {
			t.Errorf("serial number %d except %s, but %s", index, typeName, vt)
		}
	}
}

func TestVarIsPointer(t *testing.T) {
	i0 := 0
	ts := TestStruct{}
	tsPtr := &ts
	cases := map[any]bool{
		0:      false,
		&i0:    true,
		ts:     false,
		tsPtr:  true,
		&tsPtr: true,
	}
	index := 0
	for variables, except := range cases {
		index++
		result := VarIsPointer(variables)
		if result != except {
			t.Errorf("serial number %d except %v, but %v", index, except, result)
		}
	}
}

func TestName(t *testing.T) {
	test()
}
