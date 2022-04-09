/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
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
// func reverseList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}

// 	s1, s2, n := head, head.Next, head.Next.Next

// 	s1.Next = nil
// 	for s2 != nil {
// 		s2.Next = s1
// 		s1 = s2

// 		if n == nil {
// 			break
// 		}
// 		s2 = n
// 		n = n.Next
// 	}

// 	return s2
// }
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode

	for head != nil {
		head, head.Next, pre = head.Next, pre, head
		// head.Next, pre, head = pre, head, head.Next
	}

	return pre
}

// func reverseList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}

// 	ret := reverseList(head.Next)

// 	head.Next.Next = head
// 	head.Next = nil

// 	return ret
// }
// @lc code=end
