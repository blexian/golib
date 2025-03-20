package leetcode

import "testing"

func TestNumberOfGoodSubarraySplits(t *testing.T) {
	test := []int{1, 1, 0}
	res := numberOfGoodSubarraySplits(test)
	if res != 1 {
		t.Errorf("numberOfGoodSubarraySplits(%v) = %d, want 1", test, res)
	}
}
