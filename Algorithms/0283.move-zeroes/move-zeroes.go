package problem0283

func moveZeroes(nums []int) {
	for c, cur := 0, 0; cur < len(nums); cur++ {
		if nums[cur] != 0 {
			nums[c], nums[cur] = nums[cur], nums[c]
			c++
		}
	}
}
