package problem0033

func search(nums []int, target int) int {
	ro := indexOfMin(nums)
	s := len(nums)
	l, r := 0, s-1
	for l <= r {
		mid := (l + r) / 2
		rmid := (ro + mid) % s
		switch {
		case nums[rmid] < target:
			l = mid + 1
		case nums[rmid] > target:
			r = mid - 1
		default:
			return rmid
		}
	}
	return -1
}

func indexOfMin(nums []int) int {
	s := len(nums)
	l, r := 0, s-1
	for l < r {
		m := (l + r) / 2
		if nums[r] < nums[m] {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
