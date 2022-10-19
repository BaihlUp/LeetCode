/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	vec, start, res := []int{}, 0, [][]int{}
	bracktrace(nums, vec, start, &res)
	return res
}

func bracktrace(nums, vec []int, start int, res *[][]int) {
	tmp := make([]int, len(vec))
	copy(tmp, vec)
	*res = append(*res, tmp)
	for i := start; i < len(nums); i++ {
		vec = append(vec, nums[i])
		bracktrace(nums, vec, i+1, res)
		vec = vec[:len(vec)-1]
	}
}

// @lc code=end

