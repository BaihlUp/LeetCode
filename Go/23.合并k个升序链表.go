/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, start, end int) *ListNode {
	if start == end {
		return lists[start]
	}
	if start > end {
		return nil
	}
	mid := start + (end-start)/2
	return mergeTwoList(merge(lists, start, mid), merge(lists, mid+1, end))
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	tmp := &ListNode{}
	newList := tmp
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
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
	return newList.Next
}

// @lc code=end

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// 构建测试用的链表数组
	var lists []*ListNode
	for i := 0; i < 4; i++ {
		list := &ListNode{}
		curr := list
		for j := 1; j <= 5; j++ {
			curr.Val = i*5 + j
			if j < 5 {
				curr.Next = &ListNode{}
				curr = curr.Next
			}
		}
		lists = append(lists, list)
	}

	// 调用 mergeKLists 函数进行合并
	mergedList := mergeKLists(lists)

	// 输出合并后的链表结果
	printList(mergedList)
}
