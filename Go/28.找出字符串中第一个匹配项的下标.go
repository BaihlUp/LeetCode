/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 找出字符串中第一个匹配项的下标
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	for i, _ := range haystack {
		index := i
		res := -1
		for j, val := range needle {
			if index < len(haystack) && byte(val) == byte(haystack[index]) {
				index++
			} else {
				break
			}
			res = j
		}

		if res == len(needle)-1 {
			return i
		}
	}
	return -1
}

// @lc code=end

