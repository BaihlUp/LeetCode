/*
 * @lc app=leetcode.cn id=2718 lang=cpp
 *
 * [2718] 查询后矩阵的和
 */

// @lc code=start
class Solution {
public:
    long long matrixSumQueries(int n, vector<vector<int>>& queries) {
        long long res = 0;
        unordered_set<int> iset[2];
        for(int i = queries.size()-1; i >= 0; i--) {
            int type = queries[i][0], index = queries[i][1], val = queries[i][2];
            if(!iset[type].count(index)) {
                res += (n - iset[type ^ 1].size()) * val;
                iset[type].insert(index);
            }
        }
        return res;
    }
};
// @lc code=end

