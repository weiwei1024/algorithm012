package src

func numIslands(grid [][]byte) int {
	cnt := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {

			if grid[i][j] == '0' {
				continue
			}
			if grid[i][j] == '1' {
				cnt++
				clear(&grid, i, j)
			}
		}
	}
	return cnt
}

type Pos struct {
	X int
	Y int
}

//将相邻的'1'转换为'0'
func clear(grid *[][]byte, i, j int) {
	Q := make([]Pos, 0)
	Q = append(Q, Pos{i, j})
	(*grid)[i][j] = '0'
	for len(Q) != 0 {
		pos := Q[0]
		Q = Q[1:]
		x := pos.X
		y := pos.Y
		if x-1 >= 0 && (*grid)[x-1][y] == '1' {
			Q = append(Q, Pos{x - 1, y})
			(*grid)[x-1][y] = '0' //尽快标记,否则会检查很多遍
		}
		if x+1 < len(*grid) && (*grid)[x+1][y] == '1' {
			Q = append(Q, Pos{x + 1, y})
			(*grid)[x+1][y] = '0'
		}
		if y-1 >= 0 && (*grid)[x][y-1] == '1' {
			Q = append(Q, Pos{x, y - 1})
			(*grid)[x][y-1] = '0'

		}
		if y+1 < len((*grid)[0]) && (*grid)[x][y+1] == '1' {
			Q = append(Q, Pos{x, y + 1})
			(*grid)[x][y+1] = '0'
		}

	}
}
