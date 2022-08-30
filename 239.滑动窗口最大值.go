/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	res, deque := []int{}, []int{}

	push_deque := func(val int) {
		for len(deque) > 0 && deque[len(deque)-1] < val {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, val)
	}
	for idx := 0; idx < len(nums); idx++ {
		if idx < k-1 {
			push_deque(nums[idx])
			continue
		}
		push_deque(nums[idx])
		res = append(res, deque[0])
		if deque[0] == nums[idx-k+1] {
			deque = deque[1:]
		}
	}
	return res
}

// @lc code=end

