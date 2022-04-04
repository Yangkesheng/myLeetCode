/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */
package main

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB

	for a != b {
		a, b = a.Next, b.Next

		if a == nil && b == nil {
			break
		}

		if a == nil {
			a = headB
		}

		if b == nil {
			b = headA
		}
	}

	return a
}

// @lc code=end
