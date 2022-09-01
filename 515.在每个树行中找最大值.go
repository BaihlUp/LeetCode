/*
 * @lc app=leetcode.cn id=515 lang=golang
 *
 * [515] 在每个树行中找最大值
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
func largestValues(root *TreeNode) []int {
	res := []int{}
	dfs(root, &res, 1)
	return res
}

func dfs(root *TreeNode, res *[]int, lvl int) {
	if root == nil {
		return
	}
	if len(*res) < lvl {
		*res = append(*res, root.Val)
	} else if (*res)[lvl-1] < root.Val {
		(*res)[lvl-1] = root.Val
	}
	dfs(root.Left, res, lvl+1)
	dfs(root.Right, res, lvl+1)
}

// @lc code=end

