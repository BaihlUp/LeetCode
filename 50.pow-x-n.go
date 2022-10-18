/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	N := n
	if n < 0 {
		N = -1 * n
		x = 1 / x
	}
	var res float64 = 1
	for N != 0 {
		if N%2 == 1 {
			res = res * x
		}
		x = x * x
		N /= 2
	}
	return res
}

// @lc code=end

