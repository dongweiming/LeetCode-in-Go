package problem0448

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if n < 0 {
			n = -n
		}
		index := n - 1
		if nums[index] > 0 {
			nums[index] *= -1
		}
	}
	res := &[]int{}
	for i := 1; i < len(nums)+1; i++ {
		if nums[i-1] > 0 {
			*res = append(*res, i)
		}
	}
	return *res
}
