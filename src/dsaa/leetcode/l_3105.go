package leetcode

func longestMonotonicSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return 1
	}
	m := 1
	j := 0                        // 记录递增或递减子数组的第一个序号
	reduce := nums[1] < nums[0]   // 记录当前是递减子数组
	increase := nums[1] > nums[0] // 记录当前是递增子数组
	if reduce || increase {
		m = 2
	} else {
		j = 1
	}
	for i := 2; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			if m < 2 {
				m = 2
			}
			if increase {
				if i-j+1 > m {
					m = i - j + 1
				}
			} else {
				j = i - 1
			}
			reduce = false
			increase = true
		} else if nums[i] < nums[i-1] {
			if m < 2 {
				m = 2
			}
			if reduce {
				if i-j+1 > m {
					m = i - j + 1
				}
			} else {
				j = i - 1
			}
			reduce = true
			increase = false
		} else {
			reduce = false
			increase = false
			j = i
		}
	}
	return m
}
