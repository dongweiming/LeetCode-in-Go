package problem0152

func maxProduct(nums []int) int {
	max, imin, imax := -1<<32, 1, 1
	for _, i := range nums {
		if i < 0 {
			imin, imax = imax, imin
		}
		if imax*i > i {
			imax = imax * i
		} else {
			imax = i
		}
		if imin*i < i {
			imin = imin * i
		} else {
			imin = i
		}
		if imax > max {
			max = imax
		}
	}
	return max
}
