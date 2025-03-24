package leetcode

import "testing"

func TestPushDominoes(t *testing.T) {
	res := pushDominoes("RR.L")
	if res != "RR.L" {
		t.Errorf("want RR.L got %s", res)
	}
}
