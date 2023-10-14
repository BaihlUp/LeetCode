/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var flag [127]int
	res, left, right := 0, 0, 0
	for left < len(s) {
		if right < len(s) && flag[s[right]] == 0 {
			flag[s[right]]++
			right++
		} else {
			flag[s[left]]--
			left++
		}
		res = max(res, right-left)
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

