/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	left := binarySearch(nums, target, true)
	right := binarySearch(nums, target, false)
	if left <= right && right < len(nums) && nums[left] == target && target == nums[right] {
		return []int{left, right}
	}
	return []int{-1, -1}
}

func binarySearch(nums []int, target int, flag bool) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target || (nums[mid] == target && flag) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if flag {
		return left
	} else {
		return right
	}
}

// @lc code=end

