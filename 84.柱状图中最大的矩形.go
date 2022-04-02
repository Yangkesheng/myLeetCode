/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */
package main

// @lc code=start
func largestRectangleArea(heights []int) int {
	stack, max := make([]int, 0), 0
	heights = append([]int{0}, append(heights, 0)...)

	for k, v := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > v {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			w := k - (stack[len(stack)-1] + 1)

			if w*h > max {
				max = w * h
			}
		}

		stack = append(stack, k)
	}

	return max
}

// @lc code=end

// func main() {
// 	// largestRectangleArea([]int{2, 1, 5, 6, 2, 3})
// 	fmt.Println(largestRectangleArea([]int{2}))
// }
