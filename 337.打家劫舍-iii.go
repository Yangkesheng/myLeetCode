/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
 */
package main

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// @lc code=start

func rob(root *TreeNode) int {
	return rob_max(robInter(root))
}

func robInter(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	//偷；不偷
	l1, r1 := robInter(root.Left)
	l2, r2 := robInter(root.Right)

	steal := root.Val + r1 + r2
	notSteal := rob_max(l1, r1) + rob_max(l2, r2)

	return steal, notSteal
}

func rob_max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end
