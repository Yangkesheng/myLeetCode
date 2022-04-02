/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */
package main

// @lc code=start
func sortColors(nums []int) {
	red, blue := 0, 2

	redIdx, blueIdx := 0, len(nums)-1

	for i := 0; i <= blueIdx; i++ {
		if nums[i] == red {
			nums[i], nums[redIdx] = nums[redIdx], nums[i]

			redIdx++
		} else if nums[i] == blue {
			nums[i], nums[blueIdx] = nums[blueIdx], nums[i]

			blueIdx--
			i--
		}
	}
}

// @lc code=end
