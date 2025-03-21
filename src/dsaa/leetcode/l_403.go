package leetcode

import "sort"

/*
403. 青蛙过河

一只青蛙想要过河。 假定河流被等分为若干个单元格，并且在每一个单元格内都有可能放有一块石子（也有可能没有）。 青蛙可以跳上石子，但是不可以跳入水中。
给你石子的位置列表 stones（用单元格序号 升序 表示）， 请判定青蛙能否成功过河（即能否在最后一步跳至最后一块石子上）。开始时， 青蛙默认已站在第一
块石子上，并可以假定它第一步只能跳跃 1 个单位（即只能从单元格 1 跳至单元格 2 ）。
如果青蛙上一步跳跃了 k 个单位，那么它接下来的跳跃距离只能选择为 k - 1、k 或 k + 1 个单位。 另请注意，青蛙只能向前方（终点的方向）跳跃。



示例 1：
输入：stones = [0,1,3,5,6,8,12,17]
输出：true
解释：青蛙可以成功过河，按照如下方案跳跃：
跳 1 个单位到第 2 块石子, 然后跳 2 个单位到第 3 块石子, 接着 跳 2 个单位到第 4 块石子,
然后跳 3 个单位到第 6 块石子, 跳 4 个单位到第 7 块石子, 最后，跳 5 个单位到第 8 个石子（即最后一块石子）。

示例 2：
输入：stones = [0,1,2,3,4,8,9,11]
输出：false
解释：这是因为第 5 和第 6 个石子之间的间距太大，没有可选的方案供青蛙跳跃过去。


提示：
2 <= stones.length <= 2000
0 <= stones[i] <= 231 - 1
stones[0] == 0
stones 按严格升序排列
*/

func canCross(stones []int) bool {
	l := len(stones)
	dp := make([]map[int]bool, l) // dp[i][od] 表示青蛙在i时，上次跳od步数情况下 是否能跳到最后
	for i := range dp {
		dp[i] = map[int]bool{}
	}
	var dfs func(int, int) bool
	dfs = func(i, od int) (r bool) {
		if i >= l-1 {
			return true
		}
		// 已经有跳到这里的情况直接返回结果
		if res, has := dp[i][od]; has {
			return res
		}
		defer func() { dp[i][od] = r }()
		for cd := od - 1; cd <= od+1; cd++ {
			if cd < 1 {
				continue
			}
			j := sort.SearchInts(stones, cd+stones[i])
			if j != l && stones[j]-stones[i] == cd && dfs(j, cd) {
				return true
			}
		}
		return false
	}
	return dfs(0, 0)
}

// dp[i][k] = dp[j][k+1] || dp[j][k] || dp[j][k-1]
/*
i表示当前位置，k表示上次跳的步数；j表示上次的位置；那么跳到上次的步数一定是 k-1 到 k+1

*/
