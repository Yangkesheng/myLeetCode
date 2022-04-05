/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */
package main

// @lc code=start
func majorityElement(nums []int) int {
	n, count := nums[0], 0

	for _, v := range nums {
		if n == v {
			count++
		} else {
			count--

			if count < 0 {
				n = v
				count = 1
			}
		}
	}

	return n
}

// func quickSort(nums []int) {
// 	if len(nums) < 2 {
// 		return
// 	}

// 	index := partiaion(nums)

// 	if index > 1 {
// 		quickSort(nums[0:index])
// 	}

// 	if index+1 < len(nums) {
// 		quickSort(nums[index+1:])
// 	}
// }

// func partiaion(nums []int) int {
// 	start, end := 0, len(nums)-1

// 	if end-start < 2 {
// 		return start
// 	}

// 	small := start - 1
// 	for ; start < end; start++ {
// 		if nums[start] < nums[end] {

// 			small++
// 			if small != start {
// 				nums[small], nums[start] = nums[start], nums[small]
// 			}
// 		}
// 	}

// 	small++
// 	nums[small], nums[end] = nums[end], nums[small]

// 	return small
// }

// @lc code=end
// func main() {
// 	nums := []int{3, 1, 2, 5, 4, 3}
// 	quickSort(nums)

// 	fmt.Println(nums)
// }
