/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
 */
package main

import (
	"strconv"
	"strings"
)

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "NUll"
	}

	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	list := strings.Split(data, ",")

	return this.buildTree(&list)
}

func (this *Codec) buildTree(list *[]string) *TreeNode {
	now := (*list)[0]
	*list = (*list)[1:]
	if now == "NUll" {
		return nil
	}

	v, _ := strconv.Atoi(now)
	root := &TreeNode{Val: v}

	root.Left = this.buildTree(list)
	root.Right = this.buildTree(list)

	return root
}

// func buildTree(list *[]string) *TreeNode {
// 	rootVal := (*list)[0]
// 	*list = (*list)[1:]
// 	if rootVal == "X" {
// 		return nil
// 	}
// 	Val, _ := strconv.Atoi(rootVal)
// 	root := &TreeNode{Val: Val}
// 	root.Left = buildTree(list)
// 	root.Right = buildTree(list)
// 	return root
// }

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end
func main() {
	
}
