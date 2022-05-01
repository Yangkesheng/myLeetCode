/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */
package main

// @lc code=start
func countBits(n int) []int {
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = dp[i&(i-1)] + 1
	}

	return dp
}

// @lc code=end
// func main() {
// 	fmt.Println(countBits(0))
// 	fmt.Println(countBits(1))
// 	fmt.Println(countBits(2))
// 	fmt.Println(countBits(3))
// 	fmt.Println(countBits(4))
// }
