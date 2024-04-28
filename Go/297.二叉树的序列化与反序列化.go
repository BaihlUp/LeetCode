/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
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

import (
	"strings"
	"strconv"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
	res := ""
	var helper func(*TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			res = res + "null,"
		} else {
			res = res + strconv.Itoa(node.Val) + ","
			helper(node.Left)
			helper(node.Right)
		}
	}
	helper(root)
	return res
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	sp := strings.Split(data, ",")
	var helper func() *TreeNode
	helper = func() *TreeNode {
		if sp[0] == "null" {
			sp = sp[1:]
			return nil
		} else {
			val, _ := strconv.Atoi(sp[0])
			sp = sp[1:]
			return &TreeNode{val, helper(), helper()}
		}
	}
	return helper()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end

