package search

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		sortedArr []int
		value     int
		want      int
	}{
		{[]int{1, 2, 3, 4}, 1, 0},
		{[]int{1, 2, 3, 4}, 5, -1},
		{[]int{1, 2, 3, 4}, 3, 2},
	}
	for _, tst := range tests {
		res := search(tst.sortedArr, tst.value)
		if res != tst.want {
			t.Errorf("binary sort arr: %v, search: %d, want %d but got %d",
				tst.sortedArr, tst.value, tst.want, res)
		}
	}
}
