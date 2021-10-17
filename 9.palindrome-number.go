/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x <= 9 {
		return true
	} else if x%10 == 0 {
		return false
	}

	var y int
	for x > y {
		r := x % 10
		x = x / 10
		y = y * 10 + r

		if x == y || x == x / 10 == {
		return true
	}
	return false

}

// @lc code=end

