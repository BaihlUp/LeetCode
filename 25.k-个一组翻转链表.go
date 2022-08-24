/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return a
		}
		b = b.Next
	}

	newHead := reverseList(a, b)
	a.Next = reverseKGroup(b, k)
	return newHead
}

func reverseList(a, b *ListNode) *ListNode {
	pre := b

	for a != b {
		tmp := a.Next
		a.Next = pre
		pre = a
		a = tmp
	}
	return pre
}

// @lc code=end

