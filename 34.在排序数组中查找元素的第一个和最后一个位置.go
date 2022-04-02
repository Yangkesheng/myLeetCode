/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

package main

// @lc code=start
func searchRange(nums []int, target int) []int {
	pos1, pos2, left, right := -1, -1, 0, len(nums)-1

	for left <= right {
		mid := (right-left)/2 + left

		if nums[mid] == target {
			pos1 = mid
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	left, right = 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left

		if nums[mid] == target {
			pos2 = mid
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return []int{pos1, pos2}
}

// @lc code=end
