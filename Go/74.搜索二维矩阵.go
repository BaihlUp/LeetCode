/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	row, col := len(matrix), len(matrix[0])
	left, right := 0, row-1
	for left < right {
		mid := left + (right-left)/2
		if matrix[mid][col-1] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	fmt.Println("row: ", row)
	row = left
	left, right = 0, col-1
	for left <= right {
		mid := left + (right-left)/2
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// @lc code=end

