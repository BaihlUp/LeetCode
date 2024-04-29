/*
 * @lc app=leetcode.cn id=1299 lang=golang
 *
 * [1299] 将每个元素替换为右侧最大元素
 */

// @lc code=start
func replaceElements(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return arr
	}
	maxRight := -1
	for i := n - 1; i >= 0; i-- {
		tmp := arr[i]
		arr[i] = maxRight
		if tmp > maxRight {
			maxRight = tmp
		}
	}
	return arr
}

// @lc code=end

