package __datatype

import (
	"fmt"
	"testing"
)

func TestNonKey(t *testing.T) {
	m := map[string]map[string]string{}
	nonV := m["t"] // 返回空结构体
	fmt.Print(nonV["44"])
}
