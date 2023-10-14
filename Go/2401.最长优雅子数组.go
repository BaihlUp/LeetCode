/*
 * @lc app=leetcode.cn id=2401 lang=golang
 *
 * [2401] 最长优雅子数组
 */

// @lc code=start
func longestNiceSubarray(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		right := i + 1
		tmp := nums[i]
		for right < len(nums) && (tmp&nums[right] == 0) {
			tmp |= nums[right]
			right++
		}
		res = max(res, right-i)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

