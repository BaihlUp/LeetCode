/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	for k, v := range inorder {
		if v == preorder[0] {
			root.Left = buildTree(preorder[1:k+1], inorder[0:k])
			root.Right = buildTree(preorder[k+1:], inorder[k+1:])
		}
	}
	return root
}

// @lc code=end

