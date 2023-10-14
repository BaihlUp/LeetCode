/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i, v := range nums {
		if i > 0 && v == nums[i-1] {
			continue
		}
		subRes := threeSum(nums, i+1, target-v)
		for _, sli := range subRes {
			res = append(res, append(sli, v))
		}
	}
	return res
}

func threeSum(nums []int, start int, target int) [][]int {
	res := [][]int{}
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if sum < target {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

// @lc code=end

