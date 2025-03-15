package sort

import "fmt"

func Example_selectionSort() {
	tests := [][]int{
		{3, 1, 2},
		{3, 1, 2, 5, 4},
		{3},
	}
	for _, tst := range tests {
		fmt.Println(selectionSort(tst))
	}
	// OutPut:
	// [1 2 3]
	// [1 2 3 4 5]
	// [3]
}
