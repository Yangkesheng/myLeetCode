/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
 */
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
func pathSum(root *TreeNode, sum int) int {
	seen := make(map[int]int)

	seen[0] = 1
	var dfs func(root *TreeNode, pre int) int
	dfs = func(root *TreeNode, curSum int) int {
		if root == nil {
			return 0
		}

		curSum += root.Val
		res := seen[curSum-sum]

		seen[curSum]++
		res += dfs(root.Left, curSum)
		res += dfs(root.Right, curSum)
		seen[curSum]--

		return res
	}

	return dfs(root, 0)
}

/*
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	return upSum(root, 0, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func upSum(root *TreeNode, pre, sum int) int {
	if root == nil {
		return 0
	}

	cur := pre + root.Val

	n := 0
	if cur == sum {
		n++
	}

	return n + upSum(root.Left, cur, sum) + upSum(root.Right, cur, sum)
}
*/

// @lc code=end
// func main() {
// 	root := &TreeNode{Val: 0,
// 		Left:  &TreeNode{Val: 1},
// 		Right: &TreeNode{Val: 1}}

// 	fmt.Printf("%d", pathSum(root, 1))
// }
