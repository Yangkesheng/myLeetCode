/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	dummyHead, head1, head2 := &ListNode{}, list1, list2
// 	cur := dummyHead

// 	for head1 != nil && head2 != nil {
// 		if head1.Val <= head2.Val {
// 			cur.Next = head1
// 			head1 = head1.Next
// 		} else {
// 			cur.Next = head2
// 			head2 = head2.Next
// 		}

// 		cur = cur.Next
// 	}

// 	if head1 != nil {
// 		cur.Next = head1
// 	}

// 	if head2 != nil {
// 		cur.Next = head2
// 	}

// 	return dummyHead.Next
// }

// @lc code=end
