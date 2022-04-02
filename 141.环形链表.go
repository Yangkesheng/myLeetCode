/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
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
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	for fast, slow := head, head; fast != nil && slow != nil; {
		slow = slow.Next
		fast = fast.Next

		if fast == nil || slow == nil {
			break
		}
		fast = fast.Next

		if fast == slow {
			return true
		}
	}

	return false
}

// @lc code=end
