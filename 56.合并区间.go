/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	fmt.Println(intervals)
	res := [][]int{}
	for _, arr := range intervals {
		if len(res) == 0 || res[len(res)-1][1] < arr[0] {
			res = append(res, arr)
		} else {
			res[len(res)-1][1] = max(res[len(res)-1][1], arr[1])
		}
	}

	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

