/*
 * @lc app=leetcode.cn id=621 lang=golang
 *
 * [621] 任务调度器
 */

// @lc code=start
func leastInterval(tasks []byte, n int) int {
	cnt := map[byte]int{}
	for _, c := range tasks {
		cnt[c]++
	}
	maxRow, maxCol := 0, 0
	for _, v := range cnt {
		if v > maxRow {
			maxRow, maxCol = v, 1
		} else if v == maxRow {
			maxCol++
		}
	}
	if time := (maxRow-1)*(n+1) + maxCol; time > len(tasks) {
		return time
	}
	return len(tasks)
}

// @lc code=end

