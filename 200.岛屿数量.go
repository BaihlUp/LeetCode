/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	res := 0
	for i, v := range grid {
		for j, k := range v {
			if k == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}
func dfs(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) ||
		j < 0 || j >= len(grid[i]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'

	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}

// @lc code=end

