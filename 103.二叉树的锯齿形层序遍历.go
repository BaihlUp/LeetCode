/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层序遍历
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res, queue, lvl := [][]int{}, []*TreeNode{}, 0
	queue = append(queue, root)
	for queue != nil {
		ivec, q := []int{}, queue
		queue = nil
		for _, node := range q {
			ivec = append(ivec, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if lvl%2 == 1 {
			left, right := 0, len(ivec)-1
			for left < right {
				ivec[left], ivec[right] = ivec[right], ivec[left]
				left++
				right--
			}
		}
		res = append(res, ivec)
		lvl++
	}
	return res
}

// @lc code=end

