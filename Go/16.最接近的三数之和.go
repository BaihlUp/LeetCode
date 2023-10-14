/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := math.MaxInt32
	for i, v := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := v + nums[l] + nums[r]
			if abs(sum-target) < abs(res-target) {
				res = sum
			}
			if sum > target {
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				r--
			} else {
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				l++
			}
		}
	}
	return res
}
func abs(val int) int {
	if val < 0 {
		return -1 * val
	}
	return val
}

// @lc code=end

