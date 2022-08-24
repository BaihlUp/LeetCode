/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
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
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := root.Val
	max_sum(root, &res)
	return res
}

func max_sum(node *TreeNode, res *int) int {
	if node == nil {
		return 0
	}
	left := max(0, max_sum(node.Left, res))
	right := max(0, max_sum(node.Right, res))
	*res = max(*res, left+right+node.Val)
	return max(left, right) + node.Val
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

