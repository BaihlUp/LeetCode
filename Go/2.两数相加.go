/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := &ListNode{Val: 0}
	n1, n2, carry, current := 0, 0, 0, newHead
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		} else {
			n1 = 0
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		} else {
			n2 = 0
		}
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}

	return newHead.Next
}

// // 打印链表
// func printList(head *ListNode) {
// 	for head != nil {
// 		fmt.Printf("%d -> ", head.Val)
// 		head = head.Next
// 	}
// 	fmt.Println("nil")
// }

// func main() {
// 	// 示例：输入 (2 -> 4 -> 3) + (5 -> 6 -> 4)
// 	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
// 	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

// 	// 调用函数进行相加
// 	result := addTwoNumbers(l1, l2)

// 	// 打印结果链表
// 	printList(result)
// }


// @lc code=end

