package sort

func mergeSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	mid := n / 2
	return arr2Merge1(mergeSort(nums[mid:]), mergeSort(nums[:mid]))
}

// 将两个有序数组合并为一个有序数组
func arr2Merge1(a, b []int) []int {
	an, bn := len(a), len(b)
	res := make([]int, an+bn)
	var i, j int
	for i < an && j < bn {
		if a[i] < b[j] {
			res[i+j] = a[i]
			i++
		} else {
			res[i+j] = b[j]
			j++
		}
	}
	for ; i < an; i++ {
		res[i+j] = a[i]
	}
	for ; j < bn; j++ {
		res[i+j] = b[j]
	}
	return res
}
