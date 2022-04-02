/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */
package main

import (
	"sort"
)

// @lc code=start
func threeSum(nums []int) [][]int {
	res, l := make([][]int, 0), len(nums)

	sort.Ints(nums)
	for i := 0; i < l-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		for start, end := i+1, l-1; start < end; start++ {
			total := nums[i] + nums[start] + nums[end]
			if total == 0 {
				res = append(res, []int{nums[i], nums[start], nums[end]})
			} else if total > 0 {
				end--
			} else {
				break
			}
		}

	}

	return res
}

// @lc code=end
// func main() {
// 	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
// }
