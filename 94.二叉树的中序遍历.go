/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 */

package main

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// @lc code=start
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	dfs94(root, &res)
	return res
}

func dfs94(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	dfs94(root.Left, res)
	*res = append(*res, root.Val)
	dfs94(root.Right, res)
}

// @lc code=end
