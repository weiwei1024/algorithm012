package week07

//朋友圈
func findCircleNum(M [][]int) int {
	n := len(M)
	m := len(M[0])
	u := Constructor(n)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if M[i][j] == 1 {
				u.union(i, j)
			}
		}
	}
	return u.Count
}

//并查集
type UnionFind struct {
	Count  int //当前存在的集合总数
	Parent []int
}

//构建并查集
func Constructor(n int) *UnionFind {
	u := &UnionFind{
		Count:  n,
		Parent: make([]int, n),
	}
	for i := 0; i < n; i++ {
		u.Parent[i] = i
	}
	return u
}

//查找p所在集合的代表元素
func (u *UnionFind) find(p int) int {
	for p != u.Parent[p] {
		u.Parent[p] = u.Parent[u.Parent[p]]
		p = u.Parent[p]
	}
	return p
}

//合并p,q所在集合
func (u *UnionFind) union(p, q int) {
	rootP := u.find(p)
	rootQ := u.find(q)
	if rootP == rootQ {
		return
	}
	u.Parent[rootP] = rootQ
	u.Count--
}
