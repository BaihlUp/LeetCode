/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */

// @lc code=start
func maxProfit(k int, prices []int) int {
	n := len(prices)
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	for j := 0; j <= k; j++ {
		dp[0][j][0] = 0
		dp[0][j][1] = -prices[0]
	}
	for i := 1; i < n; i++ {
		for j := k; j > 0; j-- {
			// i 表示 第i天，j 交易第j次，1 当前持有股票，0 当前不持有股票，
			// dp[i][j][0] 表示不持有股票的最大利润；dp[i][j][i] 表示持有股票的最大利润
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}
	return dp[n-1][k][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

