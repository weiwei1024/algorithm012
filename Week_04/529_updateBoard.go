package src

import "strconv"

//扫雷
//从点击位置开始，做广度优先搜索
type Pair struct {
	X int
	Y int
}

var dx = []int{-1, -1, -1, 0, 0, 1, 1, 1}
var dy = []int{-1, 0, 1, -1, 1, -1, 0, 1}

func updateBoard(board [][]byte, click []int) [][]byte {
	x := click[0]
	y := click[1]
	if board[x][y] == 'M' {
		board[x][y] = 'X' //游戏结束
		return board
	}
	Q := make([]Pair, 0)
	Q = append(Q, Pair{x, y})
	board[x][y] = 'B'
	m := len(board)
	n := len(board[0])
	for len(Q) != 0 {
		x := Q[0].X
		y := Q[0].Y
		Q = Q[1:]                //出队列
		bCnt := 0                //周围地雷数
		for i := 0; i < 8; i++ { //8个方向扫地雷数
			xNew := x + dx[i]
			yNew := y + dy[i]
			if xNew >= 0 && xNew < m && yNew >= 0 && yNew < n && board[xNew][yNew] == 'M' {
				bCnt++
			}
		}
		if bCnt == 0 {
			for i := 0; i < 8; i++ { //8个方向扫地雷数
				xNew := x + dx[i]
				yNew := y + dy[i]
				if xNew >= 0 && xNew < m && yNew >= 0 && yNew < n && board[xNew][yNew] == 'E' {
					Q = append(Q, Pair{xNew, yNew})
					board[xNew][yNew] = 'B'
				}
			}
		} else {
			board[x][y] = []byte(strconv.FormatInt(int64(bCnt), 10))[0]
		}
	}
	return board
}
