package leetcode

import "testing"

func TestMaximumTotalDamage(t *testing.T) {
	tc := []int{1, 1, 3, 4}
	res := maximumTotalDamage(tc)
	if res != 6 {
		t.Errorf("maximumTotalDamage(%v) should be 6, but %d", tc, res)
	}
}
