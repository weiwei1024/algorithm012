package src

import "fmt"

var res int

func totalNQueens2(n int) int {
	res = 0
	dfs(n, 0, 0, 0, 0)
	return res
}

func dfs(n int, row int, col int, ld int, rd int) {
	if row == n {
		res++
		fmt.Println(res)
		return
	}
	// col | ld | rd将本层不可选的位置都标记出来
	// 取反,即将可选位置都标记为1,不可选标记为0
	//((1 << uint8(n)) - 1) 由于col等为int型,取反后得到的结果可能为负数,要还原为无符号数
	bits := ^(col | ld | rd) & ((1 << uint8(n)) - 1)
	for bits > 0 {
		//取最后一位为1,其他位全部标记位0
		pick := bits & (-bits)
		//col ld rd 都是对下一层不可选位置进行标记
		//col|pick 表示第n位置为1,那么下一层的第n位则不能再选
		//(ld|pick)<<1表示由于本层第n位已选,则下一层的第n-1位不能选
		//(rd|pick)>>1表示由于本层第n位已选,则下一层的第n+1位不能选
		dfs(n, row+1, col|pick, (ld|pick)<<1, (rd|pick)>>1)
		//将最后一位为1的标记为0
		bits = bits & (bits - 1)
	}
}
