/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
 */
package main

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	var next *TreeNode

	var handler func(root *TreeNode)
	handler = func(root *TreeNode) {
		if root == nil {
			return
		}

		handler(root.Right)
		handler(root.Left)
		root.Right = next
		root.Left = nil
		next = root
	}

	handler(root)
}

// @lc code=end
