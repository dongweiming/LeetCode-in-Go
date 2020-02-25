package problem0121

func maxProfit(prices []int) int {
	min, max := 1<<31, 0
	for _, p := range prices {
		if p < min {
			min = p
		} else if p-min > max {
			max = p - min
		}
	}
	return max
}
