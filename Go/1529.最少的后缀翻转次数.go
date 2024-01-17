/*
 * @lc app=leetcode.cn id=1529 lang=golang
 *
 * [1529] 最少的后缀翻转次数
 */

// 贪心算法
// 1. 根据题目 可以选择下标 i（0 <= i < n）并翻转在 闭区间 [i, n - 1] 内的所有位。则如果 i 与 i-1 不同，则一定翻转了一次
// 2. 初始化都是 0，可以定义 pre = '0'

// @lc code=start
func minFlips(target string) int {
	res := 0
	pre := '0'
	for _, s := range target {
		if s != pre {
			res++
		}
		pre = s
	}
	return res
}
// @lc code=end

