package src

var result [][]string

func totalNQueens(n int) int {
	result = [][]string{}
	trace := make([][]string, 0)
	for i := 0; i < n; i++ {
		s := make([]string, 0)
		for j := 0; j < n; j++ {
			s = append(s, ".")
		}
		trace = append(trace, s)
	}
	helper(trace, 0)

	return len(result)
}

//回溯法
func helper2(trace [][]string, row int) {
	n := len(trace)
	//到最后一行，返回结果
	if row == n {
		str := make([]string, 0)
		for i := 0; i < n; i++ {
			s := ""
			for j := 0; j < n; j++ {
				s += trace[i][j]
			}
			str = append(str, s)
		}
		result = append(result, str)
		return
	}
	//选择
	for col := 0; col < n; col++ {
		if isValid(trace, row, col) {
			//选择
			trace[row][col] = "Q"
			//下一层
			helper2(trace, row+1)
			//撤销选择
			trace[row][col] = "."
		}
	}

}

//是否有效
func isValid2(trace [][]string, row, col int) bool {
	n := len(trace)
	//找同一列
	for r := 0; r < n; r++ {
		if trace[r][col] == "Q" {
			return false
		}
	}
	//右上角
	for r, c := row-1, col+1; r >= 0 && c < n; r, c = r-1, c+1 {
		if trace[r][c] == "Q" {
			return false
		}
	}

	//左上角
	for r, c := row-1, col-1; r >= 0 && c >= 0; r, c = r-1, c+1 {
		if trace[r][c] == "Q" {
			return false
		}
	}
	return true
}
