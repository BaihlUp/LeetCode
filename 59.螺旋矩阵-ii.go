/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	left, right, top, bottom := 0, n-1, 0, n-1
	num := 1
	for num <= n*n {
		//从左往右
		if top <= bottom {
			for i := left; i <= right; i++ {
				res[left][i] = num
				num++
			}
			top++
		}
		//从上往下
		if left <= right {
			for i := top; i <= bottom; i++ {
				res[i][right] = num
				num++
			}
			right--
		}
		// 从右往左
		if top <= bottom {
			for i := right; i >= left; i-- {
				res[bottom][i] = num
				num++
			}
			bottom--
		}
		// 从下往上
		if left <= right {
			for i := bottom; i >= top; i-- {
				res[i][left] = num
				num++
			}
			left++
		}
	}
	return res
}

// @lc code=end

