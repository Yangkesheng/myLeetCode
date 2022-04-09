/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] 除自身以外数组的乘积
 */
package main

// @lc code=start
func productExceptSelf(nums []int) []int {
	l := len(nums)
	res := make([]int, len(nums))

	for m, i := 1, 0; i < l; i++ {
		res[i] = m
		m *= nums[i]
	}

	for m, i := 1, l-1; i >= 0; i-- {
		res[i] *= m
		m *= nums[i]
	}

	return res
}

// @lc code=end
