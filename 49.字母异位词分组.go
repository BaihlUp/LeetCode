/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	tmp := make(map[[127]int][]string)
	for _, str := range strs {
		cnt := [127]int{0}
		for _, v := range str {
			cnt[v-'a']++
		}
		tmp[cnt] = append(tmp[cnt], str)
	}
	res := make([][]string, 0, len(tmp))
	for _, v := range tmp {
		res = append(res, v)
	}
	return res
}

// @lc code=end

