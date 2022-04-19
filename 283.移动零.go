/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */
package main

// @lc code=start
func moveZeroes(nums []int) {
	zeroIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if zeroIndex != i {
				nums[zeroIndex], nums[i] = nums[i], nums[zeroIndex]
			}
			zeroIndex++
		}
	}
}

// @lc code=end
