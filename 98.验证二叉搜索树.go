/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 */

package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

func isValidBST(root *TreeNode) bool {
	pre := math.MinInt64

	return dfs98(root, &pre)
}

func dfs98(root *TreeNode, pre *int) bool {
	if root == nil {
		return true
	}

	if !dfs98(root.Left, pre) {
		return false
	}

	if *pre >= root.Val {
		return false
	}
	*pre = root.Val

	if !dfs98(root.Right, pre) {
		return false
	}

	return true
}

// @lc code=end
// func main() {
// 	fmt.Println(isValidBST(&TreeNode{Val: 6, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 5}}}))
// }
