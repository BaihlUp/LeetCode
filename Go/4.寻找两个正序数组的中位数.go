/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1) + len(nums2)
	if len1%2 == 1 {
		midindex := len1 / 2
		return float64(getKelement(nums1, nums2, midindex+1))
	} else {
		midindex1, midindex2 := len1/2-1, len1/2
		return float64(getKelement(nums1, nums2, midindex1+1)+getKelement(nums1, nums2, midindex2+1)) / 2.0
	}
}

func getKelement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1+k-1], nums2[index2+k-1])
		}
		mid := k / 2
		newindex1 := min(index1+mid, len(nums1)) - 1
		newindex2 := min(index2+mid, len(nums2)) - 1
		if nums1[newindex1] < nums2[newindex2] {
			k -= newindex1 - index1 + 1
			index1 = newindex1 + 1
		} else {
			k -= newindex2 - index2 + 1
			index2 = newindex2 + 1
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end

