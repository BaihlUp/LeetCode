/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	vec, res := []int{}, [][]int{}
	backtrack(nums, vec, 0, &res)
	return res
}

func backtrack(nums []int, vec []int, start int, res *[][]int) {
	tmp := make([]int, len(vec))
	copy(tmp, vec)
	*res = append(*res, tmp)
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		vec = append(vec, nums[i])
		backtrack(nums, vec, i+1, res)
		vec = vec[:len(vec)-1]
	}
}

// @lc code=end

