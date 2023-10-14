/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	mapPhone := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	res := []string{}
	backtrack(digits, "", 0, mapPhone, &res)
	return res
}

func backtrack(digits string, vec string, index int, mapPhone []string, res *[]string) {
	if index == len(digits) {
		*res = append(*res, vec)
		return
	}
	str := mapPhone[digits[index]-'0']
	for _, v := range str {
		vec += string(v)
		backtrack(digits, vec, index+1, mapPhone, res)
		vec = vec[:len(vec)-1]
	}
}

// @lc code=end

