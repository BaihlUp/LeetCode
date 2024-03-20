/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
	res := ""
	flag := 0
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		result := x + y + flag
		res = strconv.Itoa(result%10) + res
		flag = result / 10
	}
	if flag != 0 {
		res = strconv.Itoa(flag) + res
	}
	return res
}

// @lc code=end

