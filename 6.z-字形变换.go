/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	martix, down, up := make([][]byte, numRows, numRows), 0, numRows-2
	for i := 0; i != len(s); {
		if down != numRows {
			martix[down] = append(martix[down], byte(s[i]))
			i++
			down++
		} else if up > 0 {
			martix[up] = append(martix[up], byte(s[i]))
			up--
			i++
		} else {
			up = numRows - 2
			down = 0
		}
	}
	res := make([]byte, 0, len(s))
	for _, item := range martix {
		for _, v := range item {
			res = append(res, v)
		}
	}
	return string(res)
}

// @lc code=end

