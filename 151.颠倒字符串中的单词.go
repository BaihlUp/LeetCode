/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 颠倒字符串中的单词
 */

// @lc code=start
func reverseWords(s string) string {
	slowindex, fastindex := 0, 0
	b := []byte(s)
	ln := len(b)

	for ln > 0 && fastindex < ln && b[fastindex] == ' ' {
		fastindex++
	}
	for ; fastindex < ln; fastindex++ {
		if fastindex-1 > 0 && b[fastindex-1] == b[fastindex] && b[fastindex] == ' ' {
			continue
		}
		b[slowindex] = b[fastindex]
		slowindex++
	}
	if slowindex-1 > 0 && b[slowindex-1] == ' ' {
		b = b[:slowindex-1]
	} else {
		b = b[:slowindex]
	}

	ln = len(b)
	reverseByte(&b, 0, ln-1)
	i := 0
	for i < ln {
		j := i
		for j < ln && b[j] != ' ' {
			j++
		}
		reverseByte(&b, i, j-1)
		i = j
		i++
	}
	return string(b)
}

func reverseByte(b *[]byte, left, right int) {
	for left < right {
		(*b)[left], (*b)[right] = (*b)[right], (*b)[left]
		left++
		right--
	}
}

// @lc code=end

