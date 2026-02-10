package sort

// 快速排序入口函数
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	return sort(arr, 0, len(arr)-1)
}

// 递归排序函数
func sort(arr []int, low, high int) []int {
	if low < high {
		// 进行分区操作并获取基准点位置
		p := partition(arr, low, high)
		// 递归排序左半部分
		sort(arr, low, p-1)
		// 递归排序右半部分
		sort(arr, p+1, high)
	}
	return arr
}

// 分区函数
func partition(arr []int, low, high int) int {
	tmp := (low + high) / 2
	stad := arr[tmp]
	i, j := low, high
	for i < j {
		// 从左找第一个大于等于基准的元素
		for arr[i] < stad && i < tmp {
			i++
		}
		if i != tmp {
			arr[tmp] = arr[i]
			tmp = i
		}

		// 从右找第一个小于等于基准的元素
		for arr[j] > stad && j > tmp {
			j--
		}
		if j != tmp {
			arr[tmp] = arr[j]
			tmp = j
		}
		i++
		j--
	}
	arr[tmp] = stad
	return tmp
}
