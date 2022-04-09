/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
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
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	newList := revert(slow)

	for newList != nil {
		if newList.Val != head.Val {
			return false
		}

		newList = newList.Next
		head = head.Next
	}

	return true
}

func revert(head *ListNode) *ListNode {
	var pre *ListNode

	for head != nil {
		head, head.Next, pre = head.Next, pre, head
	}

	return pre
}

// @lc code=end
