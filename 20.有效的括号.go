/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	stack := make([]byte, 0)
	for _, ch := range s {
		switch ch {
		case '(', '{', '[':
			stack = append(stack, byte(ch))
		case '}', ')', ']':
			if len(stack) <= 0 {
				return false
			}
			if (ch == ']' && stack[len(stack)-1] == '[') || (ch == '}' && stack[len(stack)-1] == '{') ||
				(ch == ')' && stack[len(stack)-1] == '(') {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

// @lc code=end

