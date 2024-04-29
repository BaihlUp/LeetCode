/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}
func quickSelect(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}
	pivot := nums[left]
	i := left - 1
	j := right + 1
	for i < j {
		for i++; nums[i] < pivot; i++ {
		}
		for j--; nums[j] > pivot; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if k <= j {
		return quickSelect(nums, left, j, k)
	} else {
		return quickSelect(nums, j+1, right, k)
	}
}

// @lc code=end

