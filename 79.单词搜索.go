/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
type Pair struct {
	x, y int
}

func exist(board [][]byte, word string) bool {
	row, col := len(board), len(board[0])
	used := make([][]bool, row)
	for i, _ := range used {
		used[i] = make([]bool, col)
	}
	arr := []Pair{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		used[i][j] = true
		defer func() { used[i][j] = false }()
		for _, dir := range arr {
			if newI, newJ := i+dir.x, j+dir.y; newI >= 0 && newI < row && newJ >= 0 && newJ < col && !used[newI][newJ] {
				if check(newI, newJ, k+1) {
					return true
				}
			}
		}
		return false
	}

	for i, row := range board {
		for j, _ := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// @lc code=end

