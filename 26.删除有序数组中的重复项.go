/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	idx := 0
	for i := 1; i < len(nums); i++ {
		if nums[idx] == nums[i] {
			continue
		} else {
			idx++
			nums[idx] = nums[i]
		}
	}
	return idx + 1
}

// @lc code=end

