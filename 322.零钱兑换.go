/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */
package main

// @lc code=start
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for k := range dp {
		dp[k] = amount + 1
	}

	dp[0] = 0

	for i := 0; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}

			dp[i] = Min(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// @lc code=end
// func main() {
// 	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))
// }
