/*
 * @lc app=leetcode.cn id=792 lang=golang
 *
 * [792] 匹配子序列的单词数
 */

// @lc code=start
func numMatchingSubseq(s string, words []string) int {
	index := [26][]int{}
	for k, v := range s {
		index[v-'a'] = append(index[v-'a'], k)
	}
	res := 0
	for _, word := range words {
		i, j := 0, 0
		for _, c := range word {
			if len(index[c-'a']) == 0 {
				break
			}
			pos := left_search(index[c-'a'], j)
			if pos == -1 {
				break
			}
			j = index[c-'a'][pos] + 1
			i++
		}
		if i == len(word) {
			res++
		}
	}
	return res
}

func left_search(arr []int, target int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left == len(arr) {
		return -1
	}
	return left
}

// @lc code=end

