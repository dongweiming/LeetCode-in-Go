//package problem0046
package main

import "fmt"

func permute(nums []int) [][]int {
	res := [][]int{}
	index := 0
	path := []int{}
	used := make([]bool, len(nums))
	backtrack(nums, index, path, &res, &used)
	return res
}

func backtrack(nums []int, first int, path []int, res *[][]int, used *[]bool) {
	n := len(nums)
	if first == n {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < n; i++ {
		if !(*used)[i] {
			(*used)[i] = true
			path = append(path, nums[i])
			backtrack(nums, first+1, path, res, used)
			path = path[:len(path)-1]
			(*used)[i] = false
		}
	}
	return
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
