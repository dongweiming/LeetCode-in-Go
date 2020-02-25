package problem0015

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var res [][]int
	counter := map[int]int{}
	for _, value := range nums {
		counter[value]++
	}

	uniqNums := []int{}
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}
	sort.Ints(uniqNums)

	for i := 0; i < len(uniqNums); i++ {
		if (uniqNums[i]*3 == 0) && counter[uniqNums[i]] >= 3 {
			res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i]})
		}
		x := 0 - uniqNums[i]*2
		if counter[uniqNums[i]] > 1 && counter[x] > 0 && x > uniqNums[i] {
			res = append(res, []int{uniqNums[i], uniqNums[i], x})
		}
		for j := i + 1; j < len(uniqNums); j++ {
			if (uniqNums[j]*2+uniqNums[i] == 0) && counter[uniqNums[j]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j]})
			}
			c := 0 - uniqNums[i] - uniqNums[j]
			if c > uniqNums[j] && counter[c] > 0 {
				res = append(res, []int{uniqNums[i], uniqNums[j], c})
			}
		}
	}

	return res
}
