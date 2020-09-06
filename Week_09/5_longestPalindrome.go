package train_1

//leetcode 5 最长回文子串
//状态确定: 以规模l作为遍历,假设第i和j之间形成回文，其中i和j的规模为l
//状态递推:
//dp[i][j] 为以i和j为起点和终点的字符串是否构成回文,其中i<=j
//dp[i][j] = 1的条件：s[i]=s[j] dp[i+1][j-1]=1 其他均为0,即无法构成回文
func LongestPalindrome(s string) string {

	n := len(s)
	if n == 0 {
		return ""
	}

	dp := make([][]int, 0) //以i和j开头和结尾
	for i := 0; i < n; i++ {
		d := make([]int, 0)
		for j := 0; j < n; j++ {
			d = append(d, 0)
		}
		dp = append(dp, d)
	}
	//按照规模遍历
	for l := 1; l <= n; l++ {
		//规模为1时
		if l == 1 {
			for i := 0; i < n; i++ {
				dp[i][i] = 1
			}
			continue
		}
		//规模为2时
		if l == 2 {
			for i := 0; i < n-1; i++ {
				if s[i] == s[i+1] {
					dp[i][i+1] = 1
				} else {
					dp[i][i+1] = 0
				}
			}
			continue
		}
		//规模大于2
		for i := n - l; i >= 0; i-- {
			if s[i] == s[i+l-1] && dp[i+1][i+l-2] == 1 {
				dp[i][i+l-1] = 1
			} else {
				dp[i][i+l-1] = 0
			}
		}
	}

	//返回最长回文的起点和终点
	maxLength := 0
	st := 0
	et := n - 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i > j {
				continue
			}
			l := len(string(s[i : j+1]))
			if dp[i][j] == 1 && l > maxLength {
				st = i
				et = j + 1
				maxLength = l
			}
		}
	}

	return string(s[st:et])
}
