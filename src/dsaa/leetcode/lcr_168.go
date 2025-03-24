package leetcode

/*
给你一个整数 n ，请你找出并返回第 n 个 丑数 。
说明：丑数是只包含质因数 2、3 和/或 5 的正整数；1 是丑数。

示例 1：

输入: n = 10
输出: 12
解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
提示：

1 <= n <= 1690
*/

func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	var a, b, c int
	for i := 1; i < n; i++ {
		ma := dp[a] * 2
		mb := dp[b] * 3
		mc := dp[c] * 5
		minimum := min(ma, mb, mc)
		dp[i] = minimum
		if ma == minimum {
			a++
		}
		if mb == minimum {
			b++
		}
		if mc == minimum {
			c++
		}
	}
	return dp[n-1]
}
