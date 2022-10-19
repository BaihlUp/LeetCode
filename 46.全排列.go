/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	sort.Ints(nums)
	vec, used, res := []int{}, make([]bool, len(nums)), [][]int{}
	bracktrace(nums, vec, &used, &res)
	return res
}

func bracktrace(nums, vec []int, used *[]bool, res *[][]int) {
	if len(vec) == len(nums) {
		tmp := make([]int, len(vec))
		copy(tmp, vec)
		*res = append(*res, tmp)
	}
	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			(*used)[i] = true
			vec = append(vec, nums[i])
			bracktrace(nums, vec, used, res)
			(*used)[i] = false
			vec = vec[:len(vec)-1]
		}
	}
}

// @lc code=end

