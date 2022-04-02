/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */
package main

// @lc code=start

import "math"

// import (
// 	"fmt"
// 	"math"
// )

func lengthOfLongestSubstring(s string) int {
	max := 0.0
	for start, i := 0, 0; i < len(s); i++ {
		for n := start; n < i; n++ {
			if s[n] == s[i] {
				start = n + 1

				break
			}
		}

		max = math.Max(max, float64(i-start+1))
	}

	return int(max)
}

// @lc code=end
