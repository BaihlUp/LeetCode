/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	vec, res := []int{}, [][]int{}
	backtack(candidates, target, 0, vec, 0, &res)
	return res
}

func backtack(candidates []int, target int, sum int, vec []int, index int, res *[][]int) {
	if sum == target {
		tmp := make([]int, len(vec))
		copy(tmp, vec)
		*res = append(*res, tmp)
		return
	}
	if sum > target {
		return
	}

	for i := index; i < len(candidates); i++ {
		vec = append(vec, candidates[i])
		sum += candidates[i]
		backtack(candidates, target, sum, vec, i, res)
		sum -= candidates[i]
		vec = vec[:len(vec)-1]
	}
}

// @lc code=end

