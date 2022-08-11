/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	res := 0
	fmt.Println(reflect.TypeOf(res))

	for x != 0 {
		res = res*10 + x%10
		x = x / 10
	}
	if res > 1<<31-1 || res < -(1<<31) {
		return 0
	}
	return res
}

// @lc code=end

