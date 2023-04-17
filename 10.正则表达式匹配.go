/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
func isMatch(s string, p string) bool {

    return dp(s, 0, p, 0)
}

func dp(s string, i int, p string, j int) bool {
    slen, plen := len(s), len(p)
    if j == plen {
        return i == slen 
    }
    if i == slen {
        // s = ""
        // p = "a*b*"
        for ; j+1 < plen; j += 2 {
            if p[j+1] != '*' {
                return false
            }
        }
        if j != plen {return false}
        return true
    }

    res := false
    if s[i] == p[j] || p[j] == '.' {
        if j < plen-1 && p[j+1] == '*' {
            //*匹配0次 或 多次
            res = dp(s, i+1, p, j) || dp(s, i, p, j+2)
        } else {
            res = dp(s, i+1, p, j+1)
        }
    } else {
        if j < plen-1 && p[j+1] == '*' {
            //*匹配0次
            res = dp(s, i, p, j+2)
        } else {
            res = false
        }
    }
    return res
}


// @lc code=end

class Solution {
    public:
    // 备忘录
    vector<vector<int>> memo;

    bool isMatch(string s, string p) {
        int m = s.size(), n = p.size();
        memo = vector<vector<int>>(m, vector<int>(n, -1));
        // 指针 i，j 从索引 0 开始移动
        return dp(s, 0, p, 0);
    }

    /* 计算 p[j..] 是否匹配 s[i..] */
    bool dp(string& s, int i, string& p, int j) {
        int m = s.size(), n = p.size();
        // base case
        if (j == n) {
            return i == m;
        }
        if (i == m) {
            if ((n - j) % 2 == 1) {
                return false;
            }
            for (; j + 1 < n; j += 2) {
                if (p[j + 1] != '*') {
                    return false;
                }
            }
            return true;
        }

        // 查备忘录，防止重复计算
        if (memo[i][j] != -1) {
            return memo[i][j];
        }

        bool res = false;

        if (s[i] == p[j] || p[j] == '.') {
            if (j < n - 1 && p[j + 1] == '*') {
                res = dp(s, i, p, j + 2)
                        || dp(s, i + 1, p, j);
            } else {
                res = dp(s, i + 1, p, j + 1);
            }
        } else {
            if (j < n - 1 && p[j + 1] == '*') {
                res = dp(s, i, p, j + 2);
            } else {
                res = false;
            }
        }
        // 将当前结果记入备忘录
        memo[i][j] = res;
        return res;
    }
};
// 详细解析参见：
// https://labuladong.github.io/article/?qno=10
