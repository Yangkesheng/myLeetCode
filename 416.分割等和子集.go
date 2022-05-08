/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */
package main

// @lc code=start
func canPartition(nums []int) bool {
	sum := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum&1 == 1 {
		return false
	}

	// dp := make([][]int, n)
	// for i := 0; i < n; i++ {
	// 	dp[i] = make([]int, sum/2+1)
	// 	dp[i][0] = 1
	// }
	dp := make([]int, sum/2+1)
	dp[0] = 1

	for i := 0; i < n; i++ {
		// for j := 0; j <= sum/2; j++ {
		for j := sum / 2; j >= 0; j-- {
			// dp[i][j] = dp[i-1][j]
			if nums[i] <= j {
				// dp[i][j] = dp[i-1][j] | dp[i-1][j-nums[i]]
				dp[j] |= dp[j-nums[i]]
			}
		}
	}
	return dp[sum/2] == 1
}

// @lc code=end
// func main() {
// 	// canPartition([]int{1, 5, 11, 5})
// 	canPartition([]int{1, 1, 4})
// }
