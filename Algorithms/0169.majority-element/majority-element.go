package problem0169

func majorityElement(nums []int) int {
	count := 0
	var candidate int
	for _, n := range nums {
		if count == 0 {
			candidate = n
		}
		if candidate == n {
			count++
		} else {
			count--
		}
	}
	return candidate
}
