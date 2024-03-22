/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	low, fast := head, head
	for fast != nil && fast.Next != nil {
		low = low.Next
		fast = fast.Next.Next
		if low == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	fast = head
	for low != fast {
		low = low.Next
		fast = fast.Next
	}
	return low
}

// @lc code=end

func main() {
	head := &ListNode{1}
	tmp := head
	for i := 2; i < 10; i++ {
		tmp.Next = &ListNode{i}
		tmp = tmp.Next
	}
	fmt.Println(detectCycle(head))
	return
}