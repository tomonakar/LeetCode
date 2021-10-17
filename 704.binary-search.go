/*
 * @lc app=leetcode id=704 lang=golang
 *
 * [704] Binary Search
 */

// @lc code=start
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		mid_element := nums[mid]

		if target == mid_element {
			return mid
		} else if target > mid_element {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// @lc code=end

