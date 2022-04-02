/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

package main

// @lc code=start
// func max(a, b int) int {
// 	if a >= b {
// 		return a
// 	}

// 	return b
// }

func trap(height []int) int {
	left, right, maxLeft, maxRight, sum := 1, len(height)-2, 0, 0, 0

	for left <= right {
		if height[left-1] < height[right+1] {
			maxLeft = max(height[left-1], maxLeft)

			if maxLeft > height[left] {
				sum += maxLeft - height[left]
			}
			left++
		} else {
			maxRight = max(height[right+1], maxRight)

			if maxRight > height[right] {
				sum += maxRight - height[right]
			}
			right--
		}
	}

	return sum
}

// @lc code=end
