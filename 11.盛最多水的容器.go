/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */
package main

// @lc code=start
func maxArea(height []int) int {
	res, l, r := 0, 0, len(height)-1
	w := r

	for ; w != 0; w-- {
		h := 0
		if height[l] < height[r] {
			h = height[l]
			l++
		} else {
			h = height[r]
			r--
		}

		res = max(w*h, res)
	}

	return res
}

// @lc code=end
// func main() {
// 	fmt.Println(maxArea([]int{2, 2}))
// }
