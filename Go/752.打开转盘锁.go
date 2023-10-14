/*
 * @lc app=leetcode.cn id=752 lang=golang
 *
 * [752] 打开转盘锁
 */

// @lc code=start
func openLock(deadends []string, target string) int {
	const start = "0000"
	dead := make(map[string]bool)
	for _, val := range deadends {
		dead[val] = true
	}
	if dead[start] {
		return -1
	}
	mvOne := func(str string) (res []string) {
		tmp := []byte(str)
		for i, v := range tmp {
			tmp[i] = v - 1
			if tmp[i] < '0' {
				tmp[i] = '9'
			}
			res = append(res, string(tmp))

			tmp[i] = v + 1
			if tmp[i] > '9' {
				tmp[i] = '0'
			}
			res = append(res, string(tmp))
			tmp[i] = v
		}
		return
	}

	queue := []string{start}
	step := 0
	seed := map[string]bool{start: true}
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			if queue[i] == target {
				fmt.Println(queue[i], i)
				return step
			}
			for _, v := range mvOne(queue[i]) {
				if !dead[v] && !seed[v] {
					queue = append(queue, v)
					seed[v] = true
				}
			}

		}
		queue = queue[l:]
		step++
	}
	return -1
}

// @lc code=end

