package Week_06

//64最小路径和
func minPathSum(grid [][]int) int {
	n,m:=len(grid),len(grid[0])
	dp:=make([][]int,n)
	sum:=0
	for i:=0;i<n;i++{
		dp[i] = make([]int,m)
		sum+=grid[i][0]
		dp[i][0] = sum
	}
	sum = 0
	for j:=0;j<m;j++{
		sum+=grid[0][j]
		dp[0][j] = sum
	}
	for i:=1;i<n;i++{
		for j:=1;j<m;j++{
			dp[i][j] = min(dp[i-1][j],dp[i][j-1])+grid[i][j]
		}
	}
	return dp[n-1][m-1]
}

func min(x,y int) int{
	if x<y{
		return x
	}
	return y
}
