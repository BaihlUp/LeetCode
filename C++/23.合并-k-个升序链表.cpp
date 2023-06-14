/*
 * @lc app=leetcode.cn id=23 lang=cpp
 *
 * [23] 合并 K 个升序链表
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
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        return mergeSort(lists, 0, lists.size() - 1);
    }

    ListNode* mergeSort(vector<ListNode*>& lists, int start, int end) {
        if(start == end) return lists[start];
        if(start > end) return NULL;
        int mid = start + (end - start)/2;
        
        return mergeList(mergeSort(lists, start, mid), mergeSort(lists, mid+1, end));
    }

    ListNode* mergeList(ListNode* a, ListNode* b) {
        ListNode head, *tmp = &head, *l1 = a, *l2 = b;
        while(l1 && l2) {
            if(l1->val > l2->val) {
                tmp->next = l2;
                l2 = l2->next;
            } else {
                tmp->next = l1;
                l1 = l1->next;
            }
            tmp = tmp->next;
        }

        tmp->next = l1 ? l1 : l2;

        return head.next;
    }

};
// @lc code=end

