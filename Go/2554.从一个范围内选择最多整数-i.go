/*
 * @lc app=leetcode.cn id=2554 lang=golang
 *
 * [2554] 从一个范围内选择最多整数 I
 */

// @lc code=start
func maxCount(banned []int, n int, maxSum int) int {
	// 贪心算法
    bannedmap := make(map[int]bool)
    for _, val := range banned {
        bannedmap[val] = true
    }
    sum := 0
    cnt := 0
    for i := 1; i <= n; i++ {
        if !bannedmap[i] && sum + i <= maxSum {
            cnt++
            sum += i
        } else if sum + i > maxSum {
            break
        }
    }
    return cnt
}
// @lc code=end

