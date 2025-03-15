package sort

// 冒泡排序
func bubbleSort(nums []int) []int {
	n := len(nums)
	for i := n; i > 1; i-- {
		for j := 0; j < i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
