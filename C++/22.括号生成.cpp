/*
 * @lc app=leetcode.cn id=22 lang=cpp
 *
 * [22] 括号生成
 */

// @lc code=start
class Solution {
public:
    vector<string> res;
    vector<string> generateParenthesis(int n) {
        dfs("", 0, 0, n);
        return res;
    }
    void dfs(string istr, int left, int right, int n) {
        if(left == n && right == n){
            res.push_back(istr);
            return;
        }
            
        if(left < n)
            dfs(istr+"(", left+1, right, n);
        if(left > right && right < n)
            dfs(istr+")", left, right+1, n);
    }
};

// @lc code=end

