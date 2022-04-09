/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 */
package main

import "math"

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	max := 0.0
	dp := make([]int, len(matrix[0]))

	for _, row := range matrix {
		for k, v := range row {
			if v == '1' {
				dp[k] += 1

				h := dp[k]
				for j := k; j >= 0 && dp[j] > 0 && dp[k] >= (k-j+1); j-- {
					h = int(math.Min(float64(h), float64(dp[j])))

					realH := int(math.Min(float64(h), float64(k-j+1)))
					max = math.Max(max, float64(realH*realH))
				}
			} else {
				dp[k] = 0
			}
		}
	}

	return int(max)
}

// @lc code=end
