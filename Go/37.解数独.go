/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	solve(board)
}

func solve(board [][]byte) bool {
	for i, arr := range board {
		for j, ch := range arr {
			if ch == '.' {
				for c := '1'; c <= '9'; c++ {
					if isValid(board, i, j, byte(c)) {
						board[i][j] = byte(c)
						if solve(board) {
							return true
						} else {
							board[i][j] = '.'
						}
					}
				}
				return false
			}
		}
	}
	return true
}

func isValid(board [][]byte, row, col int, ch byte) bool {
	for i := 0; i < len(board); i++ {
		if board[row][i] == ch || board[i][col] == ch ||
			board[3*(row/3)+i/3][3*(col/3)+i%3] == ch {
			return false
		}
	}
	return true
}

// @lc code=end

