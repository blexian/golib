package others

import (
	"testing"
	"time"
)

func TestMyConcurrentMap(t *testing.T) {
	mcm := MyConcurrentMap{}
	mcm.Put(1, 1)
	v1, err1 := mcm.Get(1, time.Second)
	if err1 != nil {
		t.Errorf("get(1, 1s) want error nil, but got err: %v, val: %d", err1, v1)
	}
	go func() {
		time.Sleep(3 * time.Second)
		mcm.Put(2, 2)
	}()
	v2, err2 := mcm.Get(2, time.Second)
	if err2 == nil {
		t.Errorf("get(2, 1s) want error not nil, but got err: %v, val: %d", err2, v2)
	}
	v21, err21 := mcm.Get(2, 3*time.Second)
	if err21 != nil {
		t.Errorf("get(2, 3s) want error nil, but got err: %v, val: %d", err21, v21)
	}
}
