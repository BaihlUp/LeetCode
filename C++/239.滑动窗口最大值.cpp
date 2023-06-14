/*
 * @lc app=leetcode.cn id=239 lang=cpp
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
class Solution {
public:
    deque<int> idq;
    vector<int> maxSlidingWindow(vector<int>& nums, int k) {
        vector<int> res;
        for(int i = 0; i < nums.size(); i++) {
            if(i < k-1) {
                push_deque(nums[i]);
                continue;
            }
            push_deque(nums[i]);
            res.push_back(idq.front());
            if(nums[i-k+1] == idq.front()) {
                idq.pop_front();
            }
        }
        return res;
    }
    void push_deque(int val) {
        while(!idq.empty() && idq.back() < val) {
            idq.pop_back();
        }
        idq.push_back(val);
    }
};
// @lc code=end

