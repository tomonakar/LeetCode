/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	s := append(nums1, nums2...)
	sort.Ints(s)
	l := len(s)
	if l == 0 {
		return float64(0)
	}

	if l%2 == 0 {
		return float64(s[l/2-1]+s[l/2]) / 2
	}

	return float64(s[l/2])
}

// @lc code=end

