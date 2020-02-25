//package problem0039
package main

import "sort"
import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	path := []int{}
	index := 0
	dfs(candidates, target, index, path, &res)
	return res
}

func dfs(candidates []int, target int, index int, path []int, res *[][]int) {
	if target < 0 {
		return
	} else if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	for n := index; n < len(candidates); n++ {
		num := candidates[n]
		if num > target {
			break
		}
		path = append(path, num)
		dfs(candidates, target-num, n, path, res)
		path = path[:len(path)-1]

	}
	return
}

func main() {
	c := []int{2, 3, 7}
	fmt.Println(combinationSum(c, 18))
}
