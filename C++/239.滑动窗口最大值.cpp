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

#include <iostream>
#include <vector>

using namespace std;

int main() {
    Solution sol;

    // 测试示例 1
    vector<int> nums1 = {1, 3, -1, -3, 5, 3, 6, 7};
    int k1 = 3;
    vector<int> result1 = sol.maxSlidingWindow(nums1, k1);
    cout << "Result 1: ";
    for (int num : result1) {
        cout << num << " ";
    }
    cout << endl;

    // 测试示例 2
    vector<int> nums2 = {1, -1};
    int k2 = 1;
    vector<int> result2 = sol.maxSlidingWindow(nums2, k2);
    cout << "Result 2: ";
    for (int num : result2) {
        cout << num << " ";
    }
    cout << endl;

    // 测试示例 3
    vector<int> nums3 = {9, 11};
    int k3 = 2;
    vector<int> result3 = sol.maxSlidingWindow(nums3, k3);
    cout << "Result 3: ";
    for (int num : result3) {
        cout << num << " ";
    }
    cout << endl;

    return 0;
}