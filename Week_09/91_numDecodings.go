package train_1

//解码方法
//给定一个只包含数字的非空字符串，计算解码方法总数
//dp[i] 以第i个字符结尾的字符串的解码方法数
// 1-26个字母
// dp[i] = dp[i-1]+1 s[i-1]<3
// dp[i] = dp[i-1] s[i-1]>=3
// 0 1 不能单独解码，考虑101不能解码成AA
// 还是需要与前两项有关，这里还是跟斐波那切数列有关

// dp[i] = dp[i-1]+dp[i-2]的限制版本,第i位的dp受前两位影响
// s[i] = '0' 若s[i-1] = '1' 或 '2' dp[i] = dp[i-2]
// s[i-1] = '1' dp[i] = dp[i-1] + dp[i-2]
// s[i-1] = '2' && 1<=s[i] <= 6, dp[i] = dp[i-1] + dp[i-2]
func NumDecodings(s string) int {

	n := len(s)
	if n <= 0 {
		return 0
	}

	dp := make([]int, 0)
	for i := 0; i < n; i++ {
		dp = append(dp, 0)
	}

	for i := 0; i < n; i++ {

		if i == 0 {
			if s[i] == '0' {
				dp[i] = 0
			} else {
				dp[i] = 1
			}
			continue
		}

		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				if i-2 < 0 {
					dp[i] = 1
				} else {
					dp[i] = dp[i-2]
				}
			} else {
				return 0
			}
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] >= '1' && s[i] <= '6') {
			if i-2 < 0 {
				dp[i] = dp[i-1] + 1
			} else {
				dp[i] = dp[i-1] + dp[i-2]
			}
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[n-1]
}
