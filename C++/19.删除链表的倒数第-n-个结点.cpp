/*
 * @lc app=leetcode.cn id=19 lang=cpp
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        if(head == NULL) return head;
        ListNode* fast = head, *low = head;
        while(fast != NULL & n != 0) {
            fast = fast->next;
            n--;
        }
        if(fast == NULL) {
            return head->next;
        }

        while(fast->next != NULL) {
            low = low->next;
            fast = fast->next;
        }
        low->next = low->next->next;

        return head;
    }
};
// @lc code=end

