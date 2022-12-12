/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
func nextPermutation(nums []int) {
	size, pivot := len(nums), len(nums)-2
	for pivot >= 0 && nums[pivot] >= nums[pivot+1] {
		pivot--
	}
	if pivot >= 0 {
		i := size - 1
		for i >= 0 && nums[i] <= nums[pivot] {
			i--
		}
		nums[i], nums[pivot] = nums[pivot], nums[i]
	}

	reverse(nums, pivot+1, size-1)
}

func reverse(nums []int, start, end int) {
	for start <= end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// @lc code=end

