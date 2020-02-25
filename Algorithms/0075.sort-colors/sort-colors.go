package problem0075

func sortColors(nums []int) {
	p0, cur, p2 := 0, 0, len(nums)-1
	for cur <= p2 {
		if nums[cur] == 0 {
			nums[cur], nums[p0] = nums[p0], nums[cur]
			p0++
			cur++
		} else if nums[cur] == 2 {
			nums[cur], nums[p2] = nums[p2], nums[cur]
			p2--
		} else {
			cur++
		}
	}
}
