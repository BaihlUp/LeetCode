/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	iStack := []*TreeNode{}
	for root != nil || len(iStack) != 0 {
		for root != nil {
			iStack = append(iStack, root)
			root = root.Left
		}
		node := iStack[len(iStack)-1]
		fmt.Println(node.Val, root)
		res = append(res, node.Val)
		iStack = iStack[:len(iStack)-1]
		root = node.Right
	}

	return res
}

func inorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, res)
	*res = append(*res, root.Val)
	inorder(root.Right, res)
}

// @lc code=end

