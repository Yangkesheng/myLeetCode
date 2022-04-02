/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */
package main

// @lc code=start
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	seen := make([][]int, 0)

	for i := 0; i < m; i++ {
		seen = append(seen, make([]int, n))
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			seen[i][j] = grid[i][j]
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				seen[i][j] += seen[i][j-1]
			} else if j == 0 {
				seen[i][j] += seen[i-1][j]
			} else {
				seen[i][j] += min(seen[i-1][j], seen[i][j-1])
			}
		}
	}

	return seen[m-1][n-1]
}



// @lc code=end
// func main() {
// 	g := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}

// 	minPathSum(g)
// }
