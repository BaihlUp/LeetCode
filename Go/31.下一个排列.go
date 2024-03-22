/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
func nextPermutation(nums []int) {
	size, pivot := len(nums), len(nums)-2
	//从后往前遍历找到第一个升序对 (pivot, pivot+1)
	for pivot >= 0 && nums[pivot] >= nums[pivot+1] {
		pivot--
	}

	if pivot >= 0 {
		//找到第一个比pivot大的数与pivot交换
		i := size - 1
		for i >= 0 && nums[i] <= nums[pivot] {
			i--
		}
		nums[i], nums[pivot] = nums[pivot], nums[i]
	}

	//翻转nums[pivot+1, size-1]
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

