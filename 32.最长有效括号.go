/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */
package main

import "math"

// @lc code=start
func longestValidParentheses(s string) int {
	dp := make([]int, len(s)+1)

	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				dp[i+1] = dp[i-1] + 2
			} else if i-dp[i]-1 >= 0 && s[i-dp[i]-1] == '(' {
				dp[i+1] = dp[i] + 2 + dp[i-dp[i]-1]
			}
		}
	}

	max := 0.0
	for _, v := range dp {
		max = math.Max(max, float64(v))
	}

	return int(max)
}

// @lc code=end
