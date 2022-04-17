/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */
package main

// @lc code=start

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0, len(nums)-k+1)
	deque := make([]int, 0, len(nums))

	for i, v := range nums {
		for l := len(deque); l != 0 && nums[deque[l-1]] < v; l -= 1 {
			deque = deque[:l-1]
		}
		deque = append(deque, i)

		for deque[0] < i-k+1 {
			deque = deque[1:]
		}

		if i-k+1 < 0 {
			continue
		}

		res = append(res, nums[deque[0]])
	}

	return res
}

// func maxSlidingWindow(nums []int, k int) []int {
// 	d := deque.New()
// 	res := make([]int, 0, len(nums)-k+1)

// 	for i, v := range nums {
// 		for n, _ := d.Back(); d.Len() != 0 && nums[n.(int)] < v; {
// 			d.PopBack()

// 			n, _ = d.Back()
// 		}
// 		d.PushBack(i)

// 		for n, _ := d.Front(); n.(int) < i-k+1; {
// 			d.PopFront()

// 			n, _ = d.Front()
// 		}

// 		if i-k+1 < 0 {
// 			continue
// 		}

// 		n, _ := d.Front()
// 		res = append(res, nums[n.(int)])
// 	}

// 	return res
// }

// @lc code=end
// func main() {
// 	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
// }
