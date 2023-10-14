/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */

// @lc code=start
func multiply(num1 string, num2 string) string {
	res := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			multi := (num1[i] - '0') * (num2[j] - '0')
			sum := int(multi) + res[i+j+1]
			res[i+j+1] = sum % 10
			res[i+j] += sum / 10
		}
	}
	var str string
	flag := true
	for _, v := range res {
		if v == 0 && flag {
			continue
		}
		flag = false
		str += strconv.Itoa(v)
	}
	if len(str) == 0 {
		return "0"
	} else {
		return str
	}
}

// @lc code=end

