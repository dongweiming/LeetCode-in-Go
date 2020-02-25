package problem0011

func maxArea(height []int) int {
	rs, min, maxarea, l, r := 0, 0, 0, 0, len(height)-1
	for l < r {
		if height[l] < height[r] {
			min = height[l]
		} else {
			min = height[r]
		}
		if rs = min * (r - l); rs > maxarea {
			maxarea = rs
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxarea
}
