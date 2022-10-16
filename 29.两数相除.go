/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend == math.MinInt32 {
			return math.MaxInt32
		}
		return -1 * dividend
	}

	var a int64 = int64(dividend)
	var b int64 = int64(divisor)
	sign := 1
	if (a > 0 && b < 0) || (a < 0 && b > 0) {
		sign = -1
	}
	if a < 0 {
		a = -1 * a
	}
	if b < 0 {
		b = -1 * b
	}

	var res int64 = div(a, b)
	if sign > 0 && res > math.MaxInt32 {
		return math.MaxInt32
	}
	return sign * int(res)
}

func div(a, b int64) int64 {
	if a < b {
		return 0
	}
	var count int64 = 1
	tob := int64(b)
	for tob+tob <= a {
		//结果翻倍，快速逼近
		count = count + count
		tob = tob + tob
	}
	return count + div(a-tob, b)
}

// @lc code=end

