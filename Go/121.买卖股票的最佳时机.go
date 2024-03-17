/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	res := 0
	minValue := math.MaxInt
	for i := 0; i < len(prices); i++ {
		minValue = min(minValue, prices[i])
		res = max(res, prices[i]-minValue)
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
	if a > b {
		return a
	}
	return b
}

// @lc code=end

