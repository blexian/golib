package leetcode

/*

3186. 施咒的最大总伤害
中等
相关标签
相关企业
提示
一个魔法师有许多不同的咒语。

给你一个数组 power ，其中每个元素表示一个咒语的伤害值，可能会有多个咒语有相同的伤害值。

已知魔法师使用伤害值为 power[i] 的咒语时，他们就 不能 使用伤害为 power[i] - 2 ，power[i] - 1 ，power[i] + 1 或者 power[i] + 2 的咒语。

每个咒语最多只能被使用 一次 。

请你返回这个魔法师可以达到的伤害值之和的 最大值 。



示例 1：

输入：power = [1,1,3,4]

输出：6

解释：

可以使用咒语 0，1，3，伤害值分别为 1，1，4，总伤害值为 6 。

示例 2：

输入：power = [7,1,6,6]

输出：13

解释：

可以使用咒语 1，2，3，伤害值分别为 1，6，6，总伤害值为 13 。



提示：

1 <= power.length <= 105
1 <= power[i] <= 109

*/

func maximumTotalDamage(power []int) int64 {
	l := len(power)
	if l == 0 {
		return 0
	}
	pn := make(map[int]int)  // 存伤害为power的数量
	spower := make([]int, 0) // 存伤害的升序数组，去重
	for i := range power {
		if num, ok := pn[power[i]]; ok {
			pn[power[i]] = num + 1
		} else {
			var j int
			for ; j < len(spower); j++ {
				if spower[j] > power[i] {
					break
				}
			}
			if j == 0 {
				spower = append([]int{power[i]}, spower...)
			} else {
				spower = append(spower[:j], spower[j-1:]...) // 后移
				spower[j] = power[i]
			}
			pn[power[i]] = 1
		}
	}
	dp := make([]int64, len(spower))
	// 分两种情况 选择spower[i]或者不选，取其中最大值
	// 如果选择spower[i]，则不能选择spower[i]-1，spower[i]-2. dp[i] = dp[j] + spower[i]*pn[spower[i]]
	// j：值非spower[i]-1，spower[i]-2 的最大下标 j >= -1 dp[-1] = 0
	// 如果不选择 spower[i]，则dp[i]=dp[i-1]
	dp[0] = int64(pn[spower[0]] * spower[0])
	for i := 1; i < len(spower); i++ {
		// choice
		var dp1 int64
		j := i - 1
		for ; j >= 0; j-- {
			if spower[j] < spower[i]-2 {
				break
			}
		}
		if j == -1 {
			dp1 = int64(pn[spower[i]] * spower[i])
		} else {
			dp1 = int64(pn[spower[i]]*spower[i]) + dp[j]
		}
		// no choice
		dp2 := dp[i-1]
		dp[i] = max(dp2, dp1)
	}
	return dp[len(dp)-1]
}
