package sort

import "math"

// 选择排序 升序 从前往后遍历，找最小元素
func findMinSelectionSort(arr []int) []int {
	findMinElemIndex := func(arr []int) int {
		m := math.MaxInt
		idx := 0
		for i := range arr {
			if arr[i] < m {
				m = arr[i]
				idx = i
			}
		}
		return idx
	}
	for i := 0; i < len(arr); i++ {
		m := findMinElemIndex(arr[i:])
		arr[m+i], arr[i] = arr[i], arr[m+i]
	}
	return arr
}

// 选择排序 升序 从后往前遍历，找最大元素
func selectionSort(arr []int) []int {
	findMaxElemIndex := func(arr []int) int {
		m := 0
		for i := range arr {
			if arr[i] > arr[m] {
				m = i
			}
		}
		return m
	}
	for i := len(arr) - 1; i >= 0; i-- {
		m := findMaxElemIndex(arr[:i+1])
		arr[m], arr[i] = arr[i], arr[m]
	}
	return arr
}
