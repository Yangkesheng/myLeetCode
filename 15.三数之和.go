/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

package main

import "sort"

// @lc code=start
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)
	rtn := make([][]int, 0)

	for i := 0; i <= len(nums)-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		for start, end := i+1, len(nums)-1; start < end; start++ {
			if start > i+1 && nums[start] == nums[start-1] {
				continue
			}

			for end-1 > start && nums[i]+nums[start]+nums[end] > 0 {
				end--
			}

			if nums[i]+nums[start]+nums[end] == 0 {
				rtn = append(rtn, []int{nums[i], nums[start], nums[end]})
			}

		}
	}

	return rtn
}

// @lc code=end
