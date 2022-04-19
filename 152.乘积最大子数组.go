/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */
package main

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}

// 	return b
// }

func maxProduct(nums []int) int {
	res, maxF, minF := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		m, n := maxF, minF

		maxF = max(nums[i], max(nums[i]*m, nums[i]*n))
		minF = min(nums[i], min(nums[i]*m, nums[i]*n))

		res = max(res, maxF)
	}

	return res
}

// @lc code=end
// func main() {
// 	fmt.Println(maxProduct([]int{-2, 3, -4}))
// }
