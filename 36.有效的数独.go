/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	for i, arr := range board {
		for j, ch := range arr {
			if board[i][j] == '.' {
				continue
			}
			board[i][j] = '.'
			if !isValid(board, i, j, ch) {
				return false
			}
			board[i][j] = ch
		}
	}
	return true
}

func isValid(board [][]byte, row, col int, tmp byte) bool {
	for i := 0; i < len(board); i++ {
		if board[row][i] != '.' && board[row][i] == tmp {
			return false
		}
		if board[i][col] != '.' && board[i][col] == tmp {
			return false
		}
		if board[3*(row/3)+i/3][3*(col/3)+i%3] != '.' && board[3*(row/3)+i/3][3*(col/3)+i%3] == tmp {
			return false
		}
	}
	return true
}

// @lc code=end

