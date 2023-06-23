/*
 * @lc app=leetcode.cn id=148 lang=cpp
 *
 * [148] 排序链表
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
    ListNode* sortList(ListNode* head) {
        return sortList(head, NULL);
    }

    ListNode* sortList(ListNode *start, ListNode *end) {
        if(start == NULL) return NULL;
        if(start->next == end) {
            start->next = NULL;
            return start;
        }
        ListNode* slow = start, *fast = start;
        while(fast != end) {
            slow = slow->next;
            fast = fast->next;
            if(fast != end) {
                fast = fast->next;
            }
        }
        ListNode* mid = slow;
        return mergeList(sortList(start, mid), sortList(mid, end));
    }
    ListNode* mergeList(ListNode* l1, ListNode* l2) {
        ListNode pre;
        ListNode *tmp = &pre, *tmp1 = l1, *tmp2 = l2;
        while(tmp1 != NULL && tmp2 != NULL) {
            if(tmp1->val < tmp2->val) {
                tmp->next = tmp1;
                tmp1 = tmp1->next;
            } else {
                tmp->next = tmp2;
                tmp2 = tmp2->next;
            }
            tmp = tmp->next;
        }
        if(tmp1 == NULL) {
            tmp->next = tmp2;
        } else {
            tmp->next = tmp1;
        }
        return pre.next;
    }
};
// @lc code=end

