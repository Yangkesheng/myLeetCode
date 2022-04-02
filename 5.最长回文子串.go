/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */
package main

// @lc code=start

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		str := longestPalindromeInter(s, i, i)
		if len(str) > len(res) {
			res = str
		}

		str = longestPalindromeInter(s, i, i+1)
		if len(str) > len(res) {
			res = str
		}
	}

	return res
}

func longestPalindromeInter(s string, l, f int) string {
	for l >= 0 && f < len(s) && s[l] == s[f] {
		l--
		f++
	}

	return s[l+1 : f]
}

// @lc code=end
