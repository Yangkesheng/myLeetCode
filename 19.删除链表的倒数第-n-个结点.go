/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	num := 0
	var dp func(head *ListNode)
	dp = func(head *ListNode) {
		if head == nil {
			return
		}

		dp(head.Next)
		num++

		if num == n+1 {
			head.Next = head.Next.Next
		}
	}

	dp(head)
	if num == n {
		head = head.Next
	}

	return head
}

// @lc code=end
