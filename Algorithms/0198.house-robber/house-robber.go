package problem0198

func rob(nums []int) int {
	pre, cur := 0, 0
	for _, n := range nums {
		temp := cur
		if pre+n > cur {
			cur = pre + n
		}
		pre = temp
	}
	return cur
}
