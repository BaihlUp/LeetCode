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
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}
		res = res*10 + x%10
		x = x / 10
	}
	return res
}

// @lc code=end

