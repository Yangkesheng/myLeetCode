/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */
package main

// @lc code=start
func uniquePaths(m int, n int) int {
	seen := make([][]int, m)
	for i := 0; i < m; i++ {
		seen[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				seen[i][j] = 1
			} else {
				seen[i][j] = seen[i-1][j] + seen[i][j-1]
			}
		}
	}

	return seen[m-1][n-1]
}

// @lc code=end
// func main() {
// 	fmt.Println(uniquePaths(3, 7))
// }
