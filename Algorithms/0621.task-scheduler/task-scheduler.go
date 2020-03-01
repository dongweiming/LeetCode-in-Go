package problem0621

func leastInterval(tasks []byte, n int) int {
	res := [26]int{}
	for _, s := range tasks {
		res[s-'A']++
	}
	sort.Ints(res[:])

	var min int
	max := res[25] - 1
	idle := max * n
	for i := 24; i >= 0 && res[i] > 0; i-- {
		min = max
		if res[i] < min {
			min = res[i]
		}
		idle -= min
	}
	if idle > 0 {
		return idle + len(tasks)
	}
	return len(tasks)
}
