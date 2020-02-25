package problem0005

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); {
		tmp := helper(s, i, i)
		if len(tmp) > len(res) {
			res = tmp
		}
		if i+1 < len(s) && s[i] == s[i+1] {
			tmp = helper(s, i, i+1)
			if len(tmp) > len(res) {
				res = tmp
			}
		}
		i++
	}
	return res
}

func helper(s string, l int, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}
