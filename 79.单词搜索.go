/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */
package main

// @lc code=start
func exist(board [][]byte, word string) bool {
	m, n, seen := len(board), len(board[0]), [][]bool{}

	for i := 0; i < m; i++ {
		seen = append(seen, make([]bool, n))
	}

	var dfs func(mm, nn int, subWord string) bool
	dfs = func(mm, nn int, subWord string) bool {
		if subWord == "" {
			return true
		}

		if mm < 0 || mm >= m || nn < 0 || nn >= n || seen[mm][nn] {
			return false
		}

		if board[mm][nn] != subWord[0] {
			return false
		}

		seen[mm][nn] = true

		res := dfs(mm-1, nn, subWord[1:]) ||
			dfs(mm, nn-1, subWord[1:]) ||
			dfs(mm+1, nn, subWord[1:]) ||
			dfs(mm, nn+1, subWord[1:])

		seen[mm][nn] = false

		return res
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				if dfs(i, j, word) {
					return true
				}
			}
		}
	}

	return false
}

// @lc code=end
// func main() {
// 	exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED")
// }
