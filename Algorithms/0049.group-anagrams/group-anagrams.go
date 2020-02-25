package problem0049

import "sort"

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func groupAnagrams(strs []string) [][]string {
	mp, res := map[string][]string{}, [][]string{}
	for _, s := range strs {
		sByte := []rune(s)
		sort.Sort(sortRunes(sByte))
		a := mp[string(sByte)]
		a = append(a, s)
		mp[string(sByte)] = a
	}

	for _, v := range mp {
		res = append(res, v)
	}
	return res
}
