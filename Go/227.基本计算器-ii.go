/*
 * @lc app=leetcode.cn id=227 lang=golang
 *
 * [227] 基本计算器 II
 */

// @lc code=start
func calculate(s string) int {
	nums, num := []int{}, 0
	sign := '+'
	for i, ch := range s {
		isDigit := ch >= '0' && ch <= '9'
		if isDigit {
			num = num*10 + int(ch-'0')
		}
		if !isDigit && ch != ' ' || i == len(s)-1 {
			switch sign {
			case '+':
				nums = append(nums, num)
			case '-':
				nums = append(nums, -num)
			case '*':
				nums[len(nums)-1] = nums[len(nums)-1] * num
			case '/':
				nums[len(nums)-1] = nums[len(nums)-1] / num
			}
			sign = ch
			num = 0
		}
	}
	res := 0
	for _, v := range nums {
		res += v
	}
	return res
}

// @lc code=end

