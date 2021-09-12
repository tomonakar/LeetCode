/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */

// @lc code=start
func reverse(x int) int {
	s := 0
	for x != 0 {
		// xの剰余を取得
		t := x % 10
		// xの10の除算(ループ用)
		x = x / 10
		// 逆順に並べ替え
		s = s*10 + t
	}

	// 桁溢れ対応
	if s < -2147483648 || s > 2147483647 {
		return 0
	}
	return s
}

// @lc code=end

