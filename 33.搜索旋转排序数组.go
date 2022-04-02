/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */
package main

// @lc code=start
func search(nums []int, target int) int {
	return searchInter(nums, target, 0, len(nums)-1)
}

func searchInter(nums []int, target, l, r int) int {
	if l == r {
		if nums[l] == target {
			return l
		} else {
			return -1
		}
	}

	mid := (r-l)/2 + l
	if nums[mid] == target {
		return mid
	} else if nums[l] <= nums[mid] {
		if nums[l] <= target && target < nums[mid] {
			return searchInter(nums, target, l, mid-1)
		} else {
			return searchInter(nums, target, mid+1, r)
		}
	} else {
		if nums[mid] < target && target <= nums[r] {
			return searchInter(nums, target, mid+1, r)
		} else {
			return searchInter(nums, target, l, mid-1)
		}
	}
}

// @lc code=end
// func main() {
// 	fmt.Println(search([]int{3, 2, 1}, 1))
// }
