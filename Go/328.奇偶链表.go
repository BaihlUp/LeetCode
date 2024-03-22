/*
 * @lc app=leetcode.cn id=328 lang=golang
 *
 * [328] 奇偶链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	newHead := head.Next
	l1 := head
	l2 := newHead
	for l2 != nil && l2.Next != nil {
		l1.Next = l2.Next
		l1 = l1.Next
		l2.Next = l1.Next
		l2 = l2.Next
	}
	l1.Next = newHead

	return head
}

// @lc code=end

