package tree

func Example_breadthFirstSearch() {
	tests := [][]int{
		{0, 10, 11, 21, -1, -1, -1},
		{0, 10, 11, 20, 21, -1, 23, 30, -1, -1, 33, -1, -1, 36, 37},
	}
	for _, tst := range tests {
		breadthFirstSearch(constructor(tst))
	}
	// OutPut:
	// [ 0 10 11 21 ]
	// [ 0 10 11 20 21 23 30 33 36 37 ]
}
