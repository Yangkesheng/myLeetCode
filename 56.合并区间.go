/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */
package main

import "sort"

// @lc code=start
func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})

	res = append(res, intervals[0])

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= res[len(res)-1][1] {
			//update last range in result case
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
		} else {
			//insert new range into result case
			res = append(res, intervals[i])
		}
	}
	return res
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }

// @lc code=end
