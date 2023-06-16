/*
 * @lc app=leetcode.cn id=394 lang=cpp
 *
 * [394] 字符串解码
 */

// @lc code=start
class Solution {
public:
    string decodeString(string s) {
        stack<int> ist_num;
        stack<string> ist_str;
        int cnt = 0;
        string t;
        for(int i = 0; i < s.size(); i++) {
            if(isdigit(s[i])) {
                cnt = cnt * 10 + s[i] - '0';
            } else if (s[i] == '[') {
                ist_num.push(cnt);
                cnt = 0;
                ist_str.push(t);
                t.clear();
            } else if (s[i] == ']') {
                int num = ist_num.top(); ist_num.pop();
                for(int j = 0; j < num; j++) ist_str.top() = ist_str.top() + t;
                t = ist_str.top(); ist_str.pop();
            } else {
                t += s[i];
            }
        }
        return ist_str.empty() ? t:ist_str.top(); 
    }
};
// @lc code=end

