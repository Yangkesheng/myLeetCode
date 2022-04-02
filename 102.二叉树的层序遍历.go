/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue, res := make([]*TreeNode, 0), make([][]int, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		curLen, row := len(queue), make([]int, 0)

		for i := 0; i < curLen; i++ {
			row = append(row, queue[i].Val)

			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		res = append(res, row)
		queue = queue[curLen:]
	}

	return res
}

// @lc code=end
