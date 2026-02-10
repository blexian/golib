package leetcode

func numberOfGoodSubarraySplits(nums []int) int {
	mod := 1_000_000_007
	// i为当前 1 的下标, j 连续为1 最左边那个
	// dp[i] = dp[j] * 2^(i-j)
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	j := -1 // 记录上次遇到的1坐标
	k := -1 // 记录上次遇到的非1坐标
	for i := range nums {
		if nums[i] == 1 {
			// 第一次遇到1
			if j == -1 {
				dp[i] = 1
				j = i
				continue
			}
			// i = 1, 前面为0
			if k == i-1 {
				// 连续为0的数量
				l0 := k - j
				dp[i] = (dp[j] * (l0 + 1)) % mod
			} else {
				// i = 1, 前面为1
				// 前面连续为1的数量
				dp[i] = dp[k+1]
			}
			j = i
		} else { // i != 1, 则结果不变
			// 第一个值非1
			if k == -1 && i == 0 {
				dp[i] = 0
				k = i
				continue
			}
			dp[i] = dp[i-1]
			k = i
		}
	}
	return dp[len(nums)-1]
}
