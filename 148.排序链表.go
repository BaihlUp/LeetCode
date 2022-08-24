/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	return sortSubList(head, nil)
}

func sortSubList(start, end *ListNode) *ListNode {
	if start == nil {
		return start
	}
	if start.Next == end {
		start.Next = nil
		return start
	}
	slow, fast := start, start
	for fast != end {
		slow = slow.Next
		fast = fast.Next
		if fast != end {
			fast = fast.Next
		}
	}
	mid := slow
	return merge(sortSubList(start, mid), sortSubList(mid, end))
}

func merge(l1, l2 *ListNode) *ListNode {
	tmp := &ListNode{}
	tmp1 := tmp
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}
	if l1 == nil {
		tmp.Next = l2
	} else {
		tmp.Next = l1
	}
	return tmp1.Next
}

// @lc code=end

