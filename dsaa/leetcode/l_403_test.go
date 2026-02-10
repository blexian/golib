package leetcode

import (
	"fmt"
	"testing"
)

func TestCanCross(t *testing.T) {
	across := canCross([]int{0, 1, 3, 5, 6, 8, 12, 17})
	if !across {
		t.Errorf("want true, got false")
	}
}

func change(b *bool) {
	*b = true
}

func TestName(t *testing.T) {
	b := new(bool)
	change(b)
	fmt.Println(*b)
}
