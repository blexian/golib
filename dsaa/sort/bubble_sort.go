package sort

// 冒泡排序
func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := n; i > 1; i-- {
		for j := 0; j < i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
