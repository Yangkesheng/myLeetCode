/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
 */
package main

import "math"

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	res := math.MinInt64

	var helper func(root *TreeNode) int
	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		l := helper(root.Left)
		l = max124(l, 0)

		r := helper(root.Right)
		r = max124(r, 0)

		res = max124(root.Val+r+l, res)

		return root.Val + max124(r, l)
	}

	helper(root)

	return res
}

func max124(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end
