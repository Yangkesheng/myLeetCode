/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */
package main

// @lc code=start
func longestConsecutive(nums []int) int {
	seen, res := make(map[int]int, 0), 0

	for _, v := range nums {
		if _, ok := seen[v]; ok {
			continue
		}

		l := seen[v-1]
		r := seen[v+1]

		res = max(l+r+1, res)
		seen[v] = -1
		seen[v-l] = l + r + 1
		seen[v+r] = l + r + 1

	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end
func main() {
	longestConsecutive([]int{4, 200, 1, 3, 2})
}
