package problem0347

func topKFrequent(nums []int, k int) []int {
	res := &[]int{}
	m := make(map[int]int, len(nums))
	tmp := make([]int, len(nums)+1)
	for _, n := range nums {
		m[n]++
	}
	rm := map[int][]int{}
	for k, v := range m {
		rm[v] = append(rm[v], k)
		tmp[v] = 1
	}
	for i := len(tmp) - 1; i >= 0; i-- {
		if tmp[i] != 0 {
			for _, vv := range rm[i] {
				if k == 0 {
					return *res
				}
				*res = append(*res, vv)
				k--
			}
		}
	}
	return *res
}
