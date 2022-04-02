/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

package main

import "math"

// @lc code=start
func canJump(nums []int) bool {
	max := 0

	for k, v := range nums {
		if k > max {
			return false
		}

		max = int(math.Max(float64(v+k), float64(max)))
	}

	return true
}

// @lc code=end
