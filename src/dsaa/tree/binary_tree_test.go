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

func ExampleDepthFirstPreSearch() {
	tests := [][]int{
		{0, 10, 11, 21, -1, -1, -1},
		{0, 10, 11, 20, 21, -1, 23, 30, -1, -1, 33, -1, -1, 36, 37},
	}
	for _, tst := range tests {
		DepthFirstPreSearch(constructor(tst))
	}
	// OutPut:
	// [ 21 10 0 11 ]
	// [ 30 20 10 21 33 0 11 36 23 37 ]
}

func ExampleDepthFirstMidSearch() {
	tests := [][]int{
		{0, 10, 11, 21, -1, -1, -1},
		{0, 10, 11, 20, 21, -1, 23, 30, -1, -1, 33, -1, -1, 36, 37},
	}
	for _, tst := range tests {
		DepthFirstMidSearch(constructor(tst))
	}
	// OutPut:
	// [ 0 10 21 11 ]
	// [ 0 10 20 30 21 33 11 23 36 37 ]
}

func ExampleDepthFirstPostSearch() {
	tests := [][]int{
		{0, 10, 11, 21, -1, -1, -1},
		{0, 10, 11, 20, 21, -1, 23, 30, -1, -1, 33, -1, -1, 36, 37},
	}
	for _, tst := range tests {
		DepthFirstPostSearch(constructor(tst))
	}
	// OutPut:
	// [ 11 0 10 21 ]
	// [ 37 23 36 11 0 33 21 10 20 30 ]
}
