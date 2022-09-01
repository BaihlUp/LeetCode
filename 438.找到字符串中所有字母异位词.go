/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	sCount, pCount := [127]int{}, [127]int{}
	sLen, pLen := len(s), len(p)
	res := []int{}
	for i, v := range p {
		sCount[s[i]-'a']++
		pCount[v-'a']++
	}
	if sCount == pCount {
		res = append(res, 0)
	}
	for i, v := range s[:sLen-pLen] {
		sCount[v-'a']--
		sCount[s[i+pLen]-'a']++
		if sCount == pCount {
			res = append(res, i+1)
		}
	}
	return res
}

// @lc code=end

