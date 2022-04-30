/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */
package main

// @lc code=start
func lengthOfLIS(nums []int) int {
	tail := make([]int, 0, len(nums))

	for _, v := range nums {
		if idx := binarySearch(tail, v); idx == len(tail) {
			tail = append(tail, v)
		} else {
			tail[idx] = v
		}
	}

	return len(tail)
}

func binarySearch(tail []int, target int) int {
	left, right := 0, len(tail)

	for left < right {
		mid := left + (right-left)/2

		if tail[mid] == target {
			return mid
		} else if tail[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

// @lc code=end
// func main() {
// 	lengthOfLIS([]int{0, 1, 0, 3, 2, 3})
// }
