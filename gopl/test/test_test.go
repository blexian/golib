package test

type NumArray struct {
	preSum []int
	nums   []int
}

func Constructor(nums []int) NumArray {
	na := NumArray{}
	preSum := make([]int, len(nums))
	preSum[0] = nums[0]
	for i, j := range nums[1:] {
		preSum[i] = preSum[i-1] + j
	}
	na.preSum = preSum
	na.nums = nums
	return na
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.preSum[right] - this.preSum[left] + this.nums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
