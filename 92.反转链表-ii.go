/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right)
	}
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

var tmp *ListNode

func reverseN(head *ListNode, m int) *ListNode {
	if m == 1 {
		tmp = head.Next
		return head
	}
	newHead := reverseN(head.Next, m-1)
	head.Next.Next = head
	head.Next = tmp
	return newHead
}

// @lc code=end

