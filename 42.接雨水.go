/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	res, l_max, r_max := 0, 0, 0
	left, right := 0, len(height)-1
	for left < right {
		l_max = max(l_max, height[left])
		r_max = max(r_max, height[right])

		if l_max < r_max {
			res += l_max - height[left]
			left++
		} else {
			res += r_max - height[right]
			right--
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

