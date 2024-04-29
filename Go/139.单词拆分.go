/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
// dp[i] 标识前 i 个字符可以被 wordDict 拼接出
func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, word := range wordDict {
		wordDictSet[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

// @lc code=end

