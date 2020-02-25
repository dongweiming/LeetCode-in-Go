package problem0001

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		other := target - n
		if a, ok := m[other]; ok {
			return []int{a, i}
		}
		m[n] = i
	}
	return nil
}
