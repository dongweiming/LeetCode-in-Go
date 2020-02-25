package problem0139

func wordBreak(s string, wordDict []string) bool {
	init := make(map[int]bool)
	m := make(map[string]bool)
	for _, w := range wordDict {
		m[w] = true
	}
	init[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			o, _ := init[j]
			if ok, _ := m[s[j:i]]; ok && o {
				init[i] = true
				break
			}
		}
	}
	return init[len(s)]
}
