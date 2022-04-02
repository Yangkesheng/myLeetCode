/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */
package main

// @lc code=start
func subsets(nums []int) [][]int {
	res := [][]int{{}}

	for _, v := range nums {
		for _, ans := range res {
			res = append(res, append(append([]int{}, ans...), v))
			// res = append(res, append(ans, v)) //bug ????
		}
	}

	return res
}

// @lc code=end
