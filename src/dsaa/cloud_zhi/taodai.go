package cloud_zhi

import (
	"sort"
)

/*
给你一个二维数组[][]int{{1,1},{1,2},{2,3},{4,5}}
int[0] {1,1}表示一个宽为1，长为1的纸袋子，当存在一个宽大于1且长也大于1的纸袋子时，
例如{2,3}说明{2,3}的纸袋子可以套上{1,1}的纸袋子。给定一个二维数组，返回这些纸袋子
最多能套多少层，比如[][]int{{1,1},{1,2},{2,3},{4,5}}就最多能套3层
请用golang实现，并说下思路
*/

// Bag 定义纸袋子的结构体
type Bag struct {
	Width  int
	Height int
}

// Bags 实现 sort.Interface 接口
type Bags []Bag

func (b Bags) Len() int      { return len(b) }
func (b Bags) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b Bags) Less(i, j int) bool {
	if b[i].Width == b[j].Width {
		return b[i].Height > b[j].Height
	}
	return b[i].Width < b[j].Width
}

// 计算最多能套多少层
func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}

	// 将二维数组转换为 Bag 结构体
	bagSlice := make([]Bag, len(envelopes))
	bags := Bags(bagSlice)
	for i, e := range envelopes {
		bags[i] = Bag{Width: e[0], Height: e[1]}
	}

	// 按照宽度升序排序，宽度相同则按高度降序排序
	sort.Sort(bags)

	// 动态规划数组
	dp := make([]int, len(bags))
	for i := range dp {
		dp[i] = 1
	}

	maxLayers := 1

	// 动态规划求解最长递增子序列
	for i := 1; i < len(bags); i++ {
		for j := 0; j < i; j++ {
			if bags[i].Width > bags[j].Width && bags[i].Height > bags[j].Height {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
		if dp[i] > maxLayers {
			maxLayers = dp[i]
		}
	}

	return maxLayers
}
