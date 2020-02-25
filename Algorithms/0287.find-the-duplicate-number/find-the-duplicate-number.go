package problem0287

func findDuplicate(nums []int) int {
	a := nums[0]
	b := nums[0]
	for {
		a = nums[a]
		b = nums[nums[b]]
		if a == b {
			break
		}
	}
	x := nums[0]
	y := a
	for x != y {
		x = nums[x]
		y = nums[y]
	}
	return x
}
