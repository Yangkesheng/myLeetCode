/*
 * @lc app=leetcode.cn id=85 lang=golang
 *
 * [85] 最大矩形
 */
package main

import "math"

// @lc code=start
func maximalRectangle(matrix [][]byte) int {
	dp, max := make([]int, len(matrix[0])), 0

	for _, row := range matrix {
		for k, v := range row {
			if v == '1' {
				dp[k] += 1

				h := math.MaxInt32
				for i := k; i >= 0 && row[i] == '1'; i-- {
					if dp[i] < h {
						h = dp[i]
					}

					w := k - i + 1
					if w*h > max {
						max = w * h
					}
				}
			} else {
				dp[k] = 0
			}
		}
	}

	return max
}

// @lc code=end
