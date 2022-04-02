/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
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
func detectCycle(head *ListNode) *ListNode {
	slow, speed := head, head

	for slow != nil && speed != nil {
		speed, slow = speed.Next, slow.Next

		if speed == nil {
			break
		} else {
			speed = speed.Next
		}

		if slow == speed {
			break
		}
	}

	if slow == nil || speed == nil {
		return nil
	}

	for speed = head; speed != slow; {
		speed, slow = speed.Next, slow.Next
	}

	return slow
}

// @lc code=end
