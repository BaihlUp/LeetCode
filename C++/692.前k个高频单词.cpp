/*
 * @lc app=leetcode.cn id=692 lang=cpp
 *
 * [692] 前K个高频单词
 */

// @lc code=start
class Solution {
public:
    vector<string> topKFrequent(vector<string>& words, int k) {
        priority_queue<pair<int, string>, vector<pair<int, string>>, Solution> ique;
        unordered_map<string, int> smap;

        for(auto word : words) {
            smap[word]++;
        }
        for(auto mp : smap) {
            ique.push(make_pair(mp.second, mp.first));
            if(ique.size() > k) {
                ique.pop();
            }
        }
        vector<string> res;
        while(!ique.empty()) {
            res.insert(res.begin(), ique.top().second);
            ique.pop();
        }
        return res;
    } 

    bool operator() (pair<int, string> a, pair<int, string> b) {
        if(a.first != b.first) {
            return a.first > b.first;
        } else {
            return a.second < b.second;
        }
    }
};
// @lc code=end

