/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */
package main

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	l := len(s)
	dp, match := make([]bool, l+1), make(map[string]bool, len(wordDict))
	for _, v := range wordDict {
		match[v] = true
	}

	dp[0] = true
	for i := 1; i <= l; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && match[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[l]
}

// @lc code=end
