/*
 * @lc app=leetcode.cn id=468 lang=golang
 *
 * [468] 验证IP地址
 */

// @lc code=start
func validIPAddress(queryIP string) string {
	ips := strings.Split(queryIP, ".")
	if len(ips) == 4 {
		for _, v := range ips {
			if len(v) > 1 && v[0] == '0' {
				return "Neither"
			}
			if num, err := strconv.Atoi(v); err != nil || num > 255 {
				return "Neither"
			}
		}
		return "IPv4"
	}

	ipv6s := strings.Split(queryIP, ":")
	if len(ipv6s) == 8 {
		for _, v := range ipv6s {
			if len(v) > 4 {
				return "Neither"
			}
			if _, err := strconv.ParseUint(v, 16, 64); err != nil {
				return "Neither"
			}
		}
		return "IPv6"
	}
	return "Neither"
}

// @lc code=end

