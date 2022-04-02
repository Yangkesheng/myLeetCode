/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */
package main

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	ans := make([]int, 0)

	var dfs func(candidates []int, target int)
	dfs = func(candidates []int, target int) {
		if target == 0 {
			new := make([]int, len(ans))
			copy(new, ans)
			res = append(res, new)
		} else if target < 0 {
			return
		}

		for k, v := range candidates {
			ans = append(ans, v)
			dfs(candidates[k:], target-v)
			ans = ans[:len(ans)-1]
		}
	}

	dfs(candidates, target)
	return res
}

// @lc code=end
