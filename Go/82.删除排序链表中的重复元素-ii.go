/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	pre := &ListNode{Next: head}
	cur := pre
	flag := false
	for cur.Next != nil && cur.Next.Next != nil {
		left, right := cur.Next, cur.Next.Next
		for right != nil && left.Val == right.Val {
			left, right = right, right.Next
			flag = true
		}
		if flag == true {
			cur.Next = right
			flag = false
		} else {
			cur = left
		}
	}
	return pre.Next
}

// @lc code=end

