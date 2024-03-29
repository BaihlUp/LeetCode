/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(v int) int {
	if v < 0 {
		return -1 * v
	}
	return v
}

// @lc code=end

