/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
func isValidBST(root *TreeNode) bool {
	return helper(root, math.Inf(-1), math.Inf(1))
}

func helper(node *TreeNode, min, max float64) bool {
	if node == nil {
		return true
	}
	v := float64(node.Val)
	if v <= min || v >= max {
		return false
	}
	return helper(node.Left, min, v) && helper(node.Right, v, max)
}

// @lc code=end

