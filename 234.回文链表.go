/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	secStart := reverseList(slow.Next)
	left, right := head, secStart
	res := true
	for res && right != nil {
		if left.Val != right.Val {
			res = false
		}
		left = left.Next
		right = right.Next
	}
	//恢复链表
	slow.Next = reverseList(secStart)
	return res
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

//时间复杂度 O(n)，空间复杂度 O(n)
func isPalindrome_1(head *ListNode) bool {
	tmp := []int{}
	for ; head != nil; head = head.Next {
		tmp = append(tmp, head.Val)
	}
	for i, v := range tmp[:len(tmp)/2] {
		if v != tmp[len(tmp)-1-i] {
			return false
		}
	}
	return true
}

//时间复杂度 O(n)，空间复杂度 O(n)
func isPalindrome_2(head *ListNode) bool {
	left := head

	var recursionList func(*ListNode) bool
	recursionList = func(node *ListNode) bool {
		if node == nil {
			return true
		}
		if !recursionList(node.Next) {
			return false
		}
		if left.Val != node.Val {
			return false
		}
		left = left.Next
		return true
	}
	return recursionList(head)
}

// @lc code=end

