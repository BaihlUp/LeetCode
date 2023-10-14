/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	nums, res := []int{}, [][]int{}
	backtrack(n, nums, k, &res, 1)
	return res
}

func backtrack(n int, nums []int, k int, res *[][]int, start int) {
	if len(nums) == k {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
	}

	for i := start; i <= n; i++ {
		nums = append(nums, i)
		backtrack(n, nums, k, res, i+1)
		nums = nums[:len(nums)-1]
	}

}

// @lc code=end

