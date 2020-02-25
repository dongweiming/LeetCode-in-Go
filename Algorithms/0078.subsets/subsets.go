package problem0078

func subsets(nums []int) [][]int {
	res := &[][]int{
		[]int(nil),
	}
	for _, n := range nums {
		for _, r := range *res {
			*res = append(*res, append([]int{n}, r...))
		}
	}
	return *res
}
