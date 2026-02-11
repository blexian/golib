package sort

import (
	"fmt"
	stdsort "sort"
	"testing"
)

func TestSectionSort(t *testing.T) {
	tests := [][]int{
		{3, 1, 2},
		{3, 1, 2, 5, 4},
		{3},
	}
	for _, tst := range tests {
		res := make([]int, len(tst))
		ori := make([]int, len(tst))
		copy(ori, tst)
		copy(res, tst)
		stdsort.IntSlice(res).Sort()
		selectionSort(tst)
		tstStr := fmt.Sprint(tst)
		resStr := fmt.Sprint(res)
		if resStr != tstStr {
			t.Fatalf("selectionSort(%v) = %s, want %s", ori, tstStr, resStr)
		}
		fmt.Printf("selectionSort(%v) = %s\n", ori, tstStr)
	}
	// OutPut:
	// [1 2 3]
	// [1 2 3 4 5]
	// [3]
}

func TestDescendingSelectionSort(t *testing.T) {
	tests := [][]int{
		{3, 1, 2},
		{3, 1, 2, 5, 4},
		{3},
	}
	for _, tst := range tests {
		res := make([]int, len(tst))
		ori := make([]int, len(tst))
		copy(ori, tst)
		copy(res, tst)
		stdsort.IntSlice(res).Sort()
		findMinSelectionSort(tst)
		tstStr := fmt.Sprint(tst)
		resStr := fmt.Sprint(res)
		if resStr != tstStr {
			t.Fatalf("descendingSelectionSort(%v) = %s, want %s", ori, tstStr, resStr)
		}
		fmt.Printf("descendingSelectionSort(%v) = %s\n", ori, tstStr)
	}
	// OutPut:
	// [1 2 3]
	// [1 2 3 4 5]
	// [3]
}
