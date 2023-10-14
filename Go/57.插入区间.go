/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	left, right := newInterval[0], newInterval[1]
	flag, res := false, [][]int{}
	for _, interval := range intervals {
		if interval[0] > right {
			if !flag {
				res = append(res, []int{left, right})
				flag = true
			}
			res = append(res, interval)
		} else if interval[1] < left {
			res = append(res, interval)
		} else {
			left = min(interval[0], left)
			right = max(interval[1], right)
		}
	}
	if !flag {
		res = append(res, []int{left, right})
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

