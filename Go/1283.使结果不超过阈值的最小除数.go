/*
 * @lc app=leetcode.cn id=1283 lang=golang
 *
 * [1283] 使结果不超过阈值的最小除数
 */

// @lc code=start
func smallestDivisor(nums []int, threshold int) int {
    maxNum := 0
    for _, num := range nums {
        if maxNum < num {
            maxNum = num
        }
    }

    left, right := 1, maxNum
    for (left < right) {
        mid := left + (right - left)/2
        sum := 0
        for _, num := range nums {
            sum += int(math.Ceil(float64(num)/float64(mid)))
        }
        if sum > threshold {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}
// @lc code=end

