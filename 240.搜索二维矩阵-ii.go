/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */
package main

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j, l := 0, len(matrix[i]); j < l; j++ {
			if target == matrix[i][j] {
				return true
			} else if target < matrix[i][j] {
				l = j
			}
		}
	}

	return false
}

// @lc code=end
