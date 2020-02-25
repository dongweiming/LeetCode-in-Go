package problem0055

func canJump(nums []int) bool {
	p := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= p {
			p = i
		}
	}
	return p == 0
}
