/*
 * @lc app=leetcode.cn id=224 lang=golang
 *
 * [224] 基本计算器
 */

// @lc code=start
func calculate(s string) int {
	s1 := []byte(s)
	return helper(&s1)
}

func helper(s *[]byte) int {
	nums, num := []int{}, 0
	sign := '+'
	for len(*s) > 0 {
		ch := (*s)[0]
		*s = (*s)[1:]
		isDigit := ch >= '0' && ch <= '9'
		if isDigit {
			num = num*10 + int(ch-'0')
		}
		if ch == '(' {
			num = helper(s)
		}
		if !isDigit && ch != ' ' || len(*s) == 0 {
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
			sign = rune(ch)
			num = 0
		}
		if ch == ')' {
			break
		}
	}
	res := 0
	for _, v := range nums {
		res += v
	}
	return res
}

// @lc code=end

