/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 */

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	res := 0
	for row, v := range grid {
		for col, k := range v {
			if k == 1 {
				res = max(res, dfs(grid, row, col))
			}
		}
	}
	return res
}
func dfs(grid [][]int, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) ||
		grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	return 1 + dfs(grid, i-1, j) + dfs(grid, i+1, j) + dfs(grid, i, j-1) + dfs(grid, i, j+1)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

