/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */
package main

import (
	"math"
)

// @lc code=start
func maxSubArray(nums []int) int {
	max, now := nums[0], nums[0]

	for k, v := range nums {
		if k == 0 {
			continue
		}

		if v > now+v {
			now = v
		} else {
			now = now + v
		}

		if now > max {
			max = now
		}
	}

	return max
}

// @lc code=end

func maxSubArray1(nums []int) int {
	now, max := float64(nums[0]), float64(nums[0])

	for i := 1; i < len(nums); i++ {
		now = math.Max(float64(nums[i]), float64(nums[i])+now)

		max = math.Max(max, now)
	}

	return int(max)
}

// func main() {
// 	// fmt.Println(maxSubArray([]int{-2, -1}))
// 	fmt.Println(maxSubArray1([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
// }
