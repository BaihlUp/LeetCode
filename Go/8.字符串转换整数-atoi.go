/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	index, sign, length := 0, 1, len(s)
	res := 0
	if length <= 0 {
		return 0
	}
	for ; index < length && s[index] == ' '; index++ {
	}
	if index >= length {
		return 0
	}
	switch s[index] {
	case '+':
		index++
	case '-':
		index++
		sign = -1
	}

	for ; index < length; index++ {
		if int(s[index]-'0') < 0 || int(s[index]-'0') > 9 {
			break
		}
		res = res*10 + int(s[index]-'0')
		if sign*res < math.MinInt32 {
			return math.MinInt32
		}
		if sign*res > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return res * sign
}

// @lc code=end
