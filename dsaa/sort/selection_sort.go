package sort

func selectionSort(nums []int) []int {
	findMaxElemIndex := func(nums []int) int {
		m := 0
		for i := range nums {
			if nums[i] > nums[m] {
				m = i
			}
		}
		return m
	}
	for i := len(nums) - 1; i >= 0; i-- {
		m := findMaxElemIndex(nums[:i+1])
		nums[m], nums[i] = nums[i], nums[m]
	}
	return nums
}
