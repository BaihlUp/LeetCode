/*
 * @lc app=leetcode.cn id=49 lang=cpp
 *
 * [49] 字母异位词分组
 */

// @lc code=start
class Solution {
public:
    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        vector<vector<string>> res;
        map<string, int> smap;
        for(auto str : strs) {
            string word = str;
            sort(word.begin(), word.end());
            if(smap.find(word) == smap.end()) {
                vector<string> ivec;
                ivec.push_back(str);
                res.push_back(ivec);
                smap[word] = res.size() - 1;
            } else {
                res[smap[word]].push_back(str);
            }
        }
        return res;
    }
};
// @lc code=end

