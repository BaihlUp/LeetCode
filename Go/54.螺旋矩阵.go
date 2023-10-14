/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1
	res, nums := []int{}, len(matrix[0])*len(matrix)

	for len(res) < nums {
		if top <= bottom {
			//从左往右
			for i := left; i <= right; i++ {
				res = append(res, matrix[top][i])
			}
			top++
		}

		if left <= right {
			//从上往下
			for i := top; i <= bottom; i++ {
				res = append(res, matrix[i][right])
			}
			right--
		}

		if top <= bottom {
			//从右往左
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			//从下往上
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}
	return res
}

// @lc code=end

