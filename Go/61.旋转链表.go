/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	len := 1
	cur := head
	for cur.Next != nil {
		len++
		cur = cur.Next
	}
	if k%len == 0 {
		return head
	}
	cur.Next = head
	for i := len - k%len; i > 0; i-- {
		cur = cur.Next
	}
	tmp := cur.Next
	cur.Next = nil
	return tmp
}

// @lc code=end

