/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 */

// @lc code=start
func restoreIpAddresses(s string) []string {
	if s == "" {
		return []string{}
	}
	ip, res := []int{}, []string{}
	dfs(s, ip, 0, &res)
	return res
}

func dfs(s string, ip []int, index int, res *[]string) {
	if len(ip) == 4 {
		if index == len(s) {
			*res = append(*res, ipToString(ip))
		}
		return
	}
	if index == len(s) {
		return
	}
	if s[index] == '0' {
		ip = append(ip, 0)
		dfs(s, ip, index+1, res)
	}
	addr := 0
	for end := index; end < len(s); end++ {
		addr = addr*10 + int(s[end]-'0')
		if addr > 0 && addr <= 0xFF {
			if end == index {
				ip = append(ip, addr)
			} else {
				ip[len(ip)-1] = addr
			}
			dfs(s, ip, end+1, res)
		} else {
			break
		}
	}
}

func ipToString(ip []int) string {
	res := strconv.Itoa(ip[0])
	for i := 1; i < len(ip); i++ {
		res += "." + strconv.Itoa(ip[i])
	}
	return res
}

// @lc code=end

