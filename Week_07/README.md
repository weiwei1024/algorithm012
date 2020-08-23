# 第七周学习总结

## 知识点回顾
### 字典树
字典树（trie树）也称为前缀树，是一种树数据结构，用于检索字符串数据集中的键。尽管哈希表可以在O(1)的时间复杂度时间内寻找键值，但是无法完成以下操作：
1. 找到具有同一前缀的全部键值.
2. 按词典序枚举字符串的数据集. 

Trie 树优于哈希表的另一个理由是，随着哈希表大小增加，会出现大量的冲突，时间复杂度可能增加到O(n),其中n是插入的键的数量。
与哈希表相比，Trie 树在存储多个具有相同前缀的键时可以使用较少的空间。
此时 Trie 树只需要O(m)的时间复杂度，其中m为键长。而在平衡树中查找键值需要O(mlogn)时间复杂度。
```go
//实现字典树
type Trie struct {
	IsEnd bool
	Next  []*Trie //下一层
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		IsEnd: false,
		Next:  make([]*Trie, 26),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	curr := this
	for i := 0; i < len(word); i++ {
		if curr.Next[word[i]-'a'] == nil {
			curr.Next[word[i]-'a'] = &Trie{
				IsEnd: false,
				Next:  make([]*Trie, 26),
			}
		}
		curr = curr.Next[word[i]-'a']
	}
	curr.IsEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	curr := this
	for i := 0; i < len(word); i++ {
		if curr.Next[word[i]-'a'] == nil {
			return false
		}
		curr = curr.Next[word[i]-'a']
	}
	if curr.IsEnd == false {
		return false
	}
	return true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for i := 0; i < len(prefix); i++ {
		if curr.Next[prefix[i]-'a'] == nil {
			return false
		}
		curr = curr.Next[prefix[i]-'a']
	}
	return true
}
```

### 并查集
并查集是一种树数据结构，主要用于判断两个元素是否处于同一集合中以及对两个元素所在的集合进行合并，可以求解组团、配对等问题。
其主要定义了两种操作：
1. 确定元素属于哪一个子集。它可以被用来确定两个元素是否属于同一子集.
2. 将两个子集合并成同一个集合.

为了更加精确的定义这些方法，需要定义如何表示集合。一种常用的策略是为每个集合选定一个固定的元素，称为代表，以表示整个集合.
```go
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
```

### 回溯+剪枝
回溯法本质其实穷举，是对状态树的遍历，是最朴素的搜索算法。而在穷举过程中，通过尽可能的排除不可能到达的路径，即称为剪枝，对状态树的"剪枝"。

### 双向BFS
BFS是广度优先搜索算法，是从开始和结束位置同时开始搜索，如果能在中间相遇，说明可以搜索到。
双向BFS是典型的双向搜索技巧。

### 启发式搜索
启发式搜索是高级搜索中的一种，区别于DFS和BFS，启发式搜索在选择下一个需要搜索的节点时，通过引入估价函数的概念，即对于每个待选节点，计算其优先级。
即优先级优先搜索算法。

### AVL树
AVL树是二叉搜索树的一种，也称为平衡二叉树，其定义：
1. 左右子树的高度差小于等于1。
2. 其每一棵子树都是平衡二叉树。

### 红黑树
红黑树有以下特点：
1. 每个节点要么是红色，要么是黑色。
2. 根节点是黑节点。
3. 每个叶节点是黑色的。
4. 相邻的节点不能都为红色。
5. 从任一节点到其每个叶子节点的所有路径所包含相同数量的黑色节点。
