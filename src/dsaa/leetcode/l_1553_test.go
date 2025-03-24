package leetcode

import "testing"

func TestMinDays(t *testing.T) {
	tc := 6
	res := minDays(tc)
	if res != 3 {
		t.Errorf("minDays(%d) = %d, want 3", tc, res)
	}
}
