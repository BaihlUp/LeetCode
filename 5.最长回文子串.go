/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res = maxPalindrome(s, i, i, res)
		res = maxPalindrome(s, i, i+1, res)
	}
	return res
}

func maxPalindrome(s string, left, right int, res string) string {
	sub, flag := "", false
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
		flag = true
	}
	if flag {
		sub = s[left+1 : right]
	}
	if len(res) < len(sub) {
		return sub
	}
	return res
}

// @lc code=end

