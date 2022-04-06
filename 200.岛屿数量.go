/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */
package main

// @lc code=start
func numIslands(grid [][]byte) int {
	seen := make([][]bool, 0)

	for _, row := range grid {
		seen = append(seen, make([]bool, len(row)))
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if 0 > i ||
			i >= len(grid) ||
			j < 0 ||
			j >= len(grid[0]) ||
			grid[i][j] == '0' ||
			seen[i][j] == true {
			return
		}

		seen[i][j] = true

		dfs(i-1, j)
		dfs(i, j-1)
		dfs(i+1, j)
		dfs(i, j+1)
	}

	res := 0
	for i, row := range grid {
		for j, v := range row {
			if v == '1' && !seen[i][j] {
				dfs(i, j)
				res++
			}
		}
	}

	return res
}

// @lc code=end
