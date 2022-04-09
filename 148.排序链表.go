/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
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
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	//cut
	mid := slow.Next
	slow.Next = nil

	return merger(sortList(head), sortList(mid))
}

func merger(h1, h2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	now := dummyHead

	for h1 != nil && h2 != nil {
		if h1.Val < h2.Val {
			now.Next = h1
			h1 = h1.Next
		} else {
			now.Next = h2
			h2 = h2.Next
		}

		now = now.Next
	}

	if h1 != nil {
		now.Next = h1
	}

	if h2 != nil {
		now.Next = h2
	}

	return dummyHead.Next
}

// @lc code=end
// func main() {
// 	head := &ListNode{
// 		Val: 4,
// 		Next: &ListNode{
// 			Val: 2,
// 			Next: &ListNode{
// 				Val: 1,
// 				Next: &ListNode{
// 					Val: 3,
// 				},
// 			},
// 		},
// 	}
// 	head = sortList(head)

// 	for ; head != nil; head = head.Next {
// 		fmt.Print(head.Val, ", ")
// 	}
// }
