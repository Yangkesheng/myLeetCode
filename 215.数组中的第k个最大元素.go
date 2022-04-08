/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */
package main

// @lc code=start
func findKthLargest(nums []int, k int) int {
	return helper(nums, 0, len(nums)-1, k)
}

func helper(nums []int, start, end, k int) int {
	l := len(nums)
	n := partiaion(nums, start, end)

	if n == l-k {
		return nums[n]
	} else if l-k < n {
		return helper(nums, start, n-1, k)
	} else {
		return helper(nums, n+1, end, k)
	}
}

func partiaion(nums []int, start, end int) int {
	small := start - 1
	for i := start; i < end; i++ {
		if nums[i] < nums[end] {
			small++

			if small != i {
				nums[small], nums[i] = nums[i], nums[small]
			}
		}
	}

	small++
	nums[small], nums[end] = nums[end], nums[small]

	return small
}

// @lc code=end
// func main() {
// 	findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
// 	// findKthLargest([]int{3,2,3,1,2,4,5,5,6}, 4)

// }
