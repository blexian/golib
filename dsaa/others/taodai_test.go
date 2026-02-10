package others

import (
	"fmt"
	"testing"
)

func TestTaoDai(t *testing.T) {
	envelopes := [][]int{{1, 1}, {1, 2}, {2, 3}, {4, 5}}
	fmt.Println("最多能套的层数:", maxEnvelopes(envelopes)) // 输出: 3
}
