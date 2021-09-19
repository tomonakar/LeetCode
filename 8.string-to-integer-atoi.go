/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	var rst int
	var sign = 1
	i := 0

	// 1. 先頭の空白無視
	for i < len(s) && ' ' == s[i] {
		i++
	}

	// 2. 符号チェック
	if i < len(s) && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	// 3. 文字の読み込み (0 - 9 以外は無視)
	intMax := 1<<31 - 1
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		// 桁溢れチェック
		// -2^31 未満の場合は、-2^31にclampし、  2^31-1以上の場合は、2^31 - 1 にclampする
		if rst > intMax/10 || (rst == intMax/10 && int(s[i]-'0') > intMax%10) {
			if sign == -1 {
				return -1 << 31
			}
			return intMax
		}
		rst = rst*10 + int(s[i]-'0')
		i++
	}
	return rst * sign
}
// @lc code=end

