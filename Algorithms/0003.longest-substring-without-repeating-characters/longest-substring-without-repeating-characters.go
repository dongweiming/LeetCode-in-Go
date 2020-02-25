package problem0003

func lengthOfLongestSubstring(s string) int {
	left, maxLen := 0, 0
	lookup := [256]int{}
	for i := range lookup {
		lookup[i] = -1
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if lookup[c] >= left {
			left = lookup[c] + 1
		} else if len := i + 1 - left; len > maxLen {
			maxLen = len
		}
		lookup[c] = i
	}
	return maxLen
}
