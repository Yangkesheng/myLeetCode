/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

package main

import (
	"sort"
)

// @lc code=start
func nextPermutation(nums []int) {
	i, j := 0, 0

	for i = len(nums) - 1; i >= 0; i-- {
		if j = i - 1; j >= 0 && nums[j] < nums[i] {
			break
		}

	}

	if j != -1 {
		for i = len(nums) - 1; i >= 0; i-- {
			if nums[i] > nums[j] {
				nums[j], nums[i] = nums[i], nums[j]
				break
			}
		}

	}
	sort.Ints(nums[j+1:])
}

// @lc code=end

// func main() {
// 	nextPermutation([]int{3, 2, 1})
// 	// nextPermutation([]int{4, 2, 0, 2, 3, 2, 0})
// }
