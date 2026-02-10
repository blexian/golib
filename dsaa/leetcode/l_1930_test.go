package leetcode

import "testing"

func TestLargestOddNumber(t *testing.T) {
	testcases := map[string]string{
		"4346832": "434683",
		"434682":  "43",
		"422682":  "",
	}
	for k, v := range testcases {
		res := largestOddNumber(k)
		if res != v {
			t.Errorf("For %s, expected %s, got %s", k, v, res)
		}
	}

}
