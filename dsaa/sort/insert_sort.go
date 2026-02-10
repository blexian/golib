package sort

// 插入排序算法实现
func insertionSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i] // 当前需要插入的元素
		j := i - 1

		// 将比key大的元素后移
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key // 插入到正确位置
	}
	return arr
}
