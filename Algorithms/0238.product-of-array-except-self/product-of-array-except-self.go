package problem0238

func productExceptSelf(nums []int) []int {
	size := len(nums)
	res := make([]int, size)
	left := 1
	for i := 0; i < size; i++ {
		res[i] = left
		left *= nums[i]
	}
	right := 1
	for j := size - 1; j >= 0; j-- {
		res[j] *= right
		right *= nums[j]
	}
	return res
}
