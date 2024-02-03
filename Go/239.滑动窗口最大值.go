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

func main() {
    // 测试示例 1
    nums1 := []int{1, 3, -1, -3, 5, 3, 6, 7}
    k1 := 3
    result1 := maxSlidingWindow(nums1, k1)
    fmt.Println("Result 1:", result1)

    // 测试示例 2
    nums2 := []int{1, -1}
    k2 := 1
    result2 := maxSlidingWindow(nums2, k2)
    fmt.Println("Result 2:", result2)

    // 测试示例 3
    nums3 := []int{9, 11}
    k3 := 2
    result3 := maxSlidingWindow(nums3, k3)
    fmt.Println("Result 3:", result3)
}