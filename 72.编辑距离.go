/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */
package main

// @lc code=start
func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)

	dp := make([][]int, 0)
	for i := 0; i < l1+1; i++ {
		row := make([]int, l2+1)
		for j := 0; j < l2+1; j++ {
			if i == 0 {
				row[j] = j
			} else if j == 0 {
				row[j] = i
			}
		}

		dp = append(dp, row)
	}

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min(min(dp[i+1][j], dp[i][j+1]), dp[i][j]) + 1
			}
		}

	}

	return dp[l1][l2]
}

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}

// 	return b
// }

// @lc code=end
// func main() {
// 	minDistance("horse", "ros")
// 	// minDistance1("horse", "ros")
// }
