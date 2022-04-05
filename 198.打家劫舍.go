/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */
package main

import "math"

// @lc code=start
func rob(nums []int) int {
	max, pre := 0.0, 0.0

	for _, v := range nums {
		preMax := max
		max = math.Max(pre+float64(+v), float64(max))

		pre = preMax
	}

	return int(max)
}

// @lc code=end
// func main() {
// 	rob([]int{1, 3, 1})
// }
