/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
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
			if i > 0 && nums[i] == nums[i-1] && !(*used)[i-1] {
				continue
			}
			(*used)[i] = true
			vec = append(vec, nums[i])
			bracktrace(nums, vec, used, res)
			(*used)[i] = false
			vec = vec[:len(vec)-1]
		}
	}
}

// @lc code=end

