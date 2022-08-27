/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	res := []string{}
	sub := ""
	backtrack(n, sub, &res, 0, 0)
	return res
}

func backtrack(n int, sub string, res *[]string, left, right int) {
	if right == n && left == n {
		*res = append(*res, sub)
		return
	}

	if left < n {
		backtrack(n, sub+"(", res, left+1, right)
	}
	if left > right && right < n {
		backtrack(n, sub+")", res, left, right+1)
	}
}

// @lc code=end

