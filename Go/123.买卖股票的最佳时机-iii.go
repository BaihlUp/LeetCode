/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */

// @lc code=start
func maxProfit(prices []int) int {
	max_k := 2
	n := len(prices)
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, max_k+1)
		for j := 0; j <= max_k; j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	for k := 1; k <= max_k; k++ {
		dp[0][k][0] = 0
		dp[0][k][1] = -prices[0]
	}
	for i := 1; i < n; i++ {
		for k := max_k; k >= 1; k-- {
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	return dp[n-1][max_k][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

