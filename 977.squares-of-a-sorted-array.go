/*
 * @lc app=leetcode id=977 lang=golang
 *
 * [977] Squares of a Sorted Array
 */

// @lc code=start
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	left, right := 0, n-1

	for i := n - 1; i >= 0; i-- {
		maxValue := 0

		if Abs(nums[left]) > Abs(nums[right]) {
			maxValue = nums[left]
			left++
		} else {
			maxValue = nums[right]
			right--
		}
		result = maxValue * maxValue
	}
	return result

}

// @lc code=end

