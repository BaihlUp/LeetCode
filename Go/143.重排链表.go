/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	//寻找中点
	midNode := midList(head)
	l1 := head
	l2 := midNode.Next
	midNode.Next = nil
	//链表逆序
	l2 = reverseList(l2)
	//合并链表
	mergeList(l1, l2)
}

func midList(root *ListNode) *ListNode {
	low, fast := root, root
	for fast != nil && fast.Next != nil {
		low = low.Next
		fast = fast.Next.Next
	}
	return low
}

func reverseList(root *ListNode) *ListNode {
	if root == nil {
		return nil
	}
	var pre *ListNode = nil
	for root != nil {
		tmp := root.Next
		root.Next = pre
		pre = root
		root = tmp
	}
	return pre
}

func mergeList(l1, l2 *ListNode) {
	var l1Tmp, l2Tmp *ListNode
	for l1 != nil && l2 != nil {
		l1Tmp = l1.Next
		l2Tmp = l2.Next

		l1.Next = l2
		l1 = l1Tmp

		l2.Next = l1
		l2 = l2Tmp
	}
}

// @lc code=end

