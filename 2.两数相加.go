/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	now := &ListNode{}
	res := now

	for e := 0; l1 != nil || l2 != nil || e != 0; {
		val := e
		if l1 != nil {
			val += l1.Val

			l1 = l1.Next
		}

		if l2 != nil {
			val += l2.Val

			l2 = l2.Next
		}

		new := &ListNode{
			Val: val % 10,
		}

		e = val / 10
		now.Next = new

		now = now.Next
	}

	return res.Next
}

// @lc code=end
