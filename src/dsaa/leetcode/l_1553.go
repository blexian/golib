package leetcode

import "math"

/*


代码
测试用例
测试用例
测试结果
1553. 吃掉 N 个橘子的最少天数
困难
相关标签
相关企业
提示
厨房里总共有 n 个橘子，你决定每一天选择如下方式之一吃这些橘子：

吃掉一个橘子。
如果剩余橘子数 n 能被 2 整除，那么你可以吃掉 n/2 个橘子。
如果剩余橘子数 n 能被 3 整除，那么你可以吃掉 2*(n/3) 个橘子。
每天你只能从以上 3 种方案中选择一种方案。

请你返回吃掉所有 n 个橘子的最少天数。



示例 1：

输入：n = 10
输出：4
解释：你总共有 10 个橘子。
第 1 天：吃 1 个橘子，剩余橘子数 10 - 1 = 9。
第 2 天：吃 6 个橘子，剩余橘子数 9 - 2*(9/3) = 9 - 6 = 3。（9 可以被 3 整除）
第 3 天：吃 2 个橘子，剩余橘子数 3 - 2*(3/3) = 3 - 2 = 1。
第 4 天：吃掉最后 1 个橘子，剩余橘子数 1 - 1 = 0。
你需要至少 4 天吃掉 10 个橘子。
示例 2：

输入：n = 6
输出：3
解释：你总共有 6 个橘子。
第 1 天：吃 3 个橘子，剩余橘子数 6 - 6/2 = 6 - 3 = 3。（6 可以被 2 整除）
第 2 天：吃 2 个橘子，剩余橘子数 3 - 2*(3/3) = 3 - 2 = 1。（3 可以被 3 整除）
第 3 天：吃掉剩余 1 个橘子，剩余橘子数 1 - 1 = 0。
你至少需要 3 天吃掉 6 个橘子。
示例 3：

输入：n = 1
输出：1
示例 4：

输入：n = 56
输出：6


提示：

1 <= n <= 2*10^9

*/

func minDays1(n int) int {
	if n < 1 {
		return 0
	}
	// 分3种情况，取最小值
	// 1、吃1个橘子 dp[n] = dp[n-1] + 1
	// 2、吃n/2 dp[n] = dp[n/2] + 1
	// 3、吃 2 * n / 3   dp[n] = dp[n/3] + 1
	dp := make([]int, n+1)
	dp[0] = 0
	maxInt := math.MaxInt
	for i := 3; i <= n; i++ {
		dp1 := dp[i-1] + 1
		dp2 := maxInt
		if i%2 == 0 {
			dp2 = dp[i/2] + 1
		}
		dp3 := maxInt
		if i%3 == 0 {
			dp3 = dp[i/3] + 1
		}
		dp[i] = min(dp1, dp2, dp3)
	}
	return dp[n]
}

func minDays(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 2
	}
	dp := make(map[int]int)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	dp[3] = 2
	return _minDays(n, dp)
}

func _minDays(n int, dp map[int]int) int {
	if v, ok := dp[n]; ok {
		return v
	}
	// 分3种情况，取最小值
	// 1、吃1个橘子 dp[n] = dp[n-1] + 1
	// 2、吃n/2 dp[n] = dp[n/2] + 1
	// 3、吃 2 * n / 3   dp[n] = dp[n/3] + 1
	dp1 := math.MaxInt
	if dp[n-1] != 0 {
		dp1 = dp[n-1] + 1
	} else {
		dp1 = _minDays(n-1, dp) + 1
	}
	dp2 := math.MaxInt
	if n%2 == 0 {
		if dp[n/2] != 0 {
			dp2 = dp[n/2] + 1
		} else {
			dp2 = _minDays(n/2, dp) + 1
		}
	}
	dp3 := math.MaxInt
	if n%3 == 0 {
		if dp[n/3] != 0 {
			dp3 = dp[n/3] + 1
		} else {
			dp3 = _minDays(n/3, dp) + 1
		}
	}
	dp[n] = min(dp1, dp2, dp3)
	return dp[n]
}
