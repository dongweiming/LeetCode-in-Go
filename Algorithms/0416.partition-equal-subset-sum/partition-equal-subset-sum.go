package problem0416

// 01背包问题
func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum&1 == 1 {
		return false
	}
	sum >>= 1
	dp := make([]bool, sum+1)
	dp[0] = true
	for _, n := range nums {
		for i := sum; i >= n; i-- {
			dp[i] = dp[i] || dp[i-n]
		}
	}
	return dp[sum]
}
