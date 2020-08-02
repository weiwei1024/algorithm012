# 第四周学习总结

## 知识点回顾
### 深度优先搜索
深度优先搜索算法（英语：Depth-First-Search，DFS）是一种用于遍历或搜索树或图的算法。沿着树的深度遍历树的节点，尽可能深的搜索树的分支。当节点v的所在边都己被探寻过，搜索将回溯到发现节点v的那条边的起始节点。这一过程一直进行到已发现从源节点可达的所有节点为止。如果还存在未被发现的节点，则选择其中一个作为源节点并重复以上过程，整个进程反复进行直到所有节点都被访问为止。属于盲目搜索。
``` go
//DFS代码模版
func dfs(node *TreeNode, level int, visited map[int]bool, result *[]*TreeNode) {
	if node == nil {
		return
	}
	//判断节点是否访问过
	if _, ok := visited[node.Val]; ok {
		return
	}
	//添加结构
	*result = append(*result, node)
	//访问左节点
	if node.Left != nil {
		dfs(node.Left, level+1, visited, result)
	}
	//访问右节点
	if node.Right != nil {
		dfs(node.Right, level+1, visited, result)
	}
}
```
### 广度优先搜索
广度优先搜索算法（英语：Breadth-First-Search，BFS）是一种用于遍历或搜索树或图的算法.可以认为BFS是从根节点开始,首先对其所有子节点遍历,之后再对其所有子节点的孙节点遍历.当所有节点被遍历后,算法终止.

``` go
//BFS代码模版
func bfs(node *TreeNode) {
	if node == nil {
		return
	}
	Q := make([]*TreeNode, 0) //构建队列
	Q = append(Q, node)
	for len(Q) != 0 {
		q := Q[0]
		Q = Q[1:] //出队列
		if q.Left != nil {
			Q = append(Q, q.Left)
		}
		if q.Right != nil {
			Q = append(Q, q.Right)
		}
	}
}
```

### 贪心算法
贪心算法在求解问题时,总是做出在当时看起来最好的选择,即不从整体上考虑最优解.如果一个问题可以通过局部最优解获取到整体最优解,则适用于贪心算法.

### 二分查找
二分查找算法可以使用的前提条件:(1)待查找的数列具有单调性(2)存在边界(3)可通过下标索引
``` go
//二分查找代码模版
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
	  //计算中间位置
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
```
## 解题思路
### 深度(广度)优先搜索
DFS和BFS都是图搜索算法,因此对于能使用这两种搜索算法的,往往可以把题目中的`状态`抽象为图中的一个节点,`状态`之间可转换,即表示节点之间可达.如[`433. 最小基因变化`](https://leetcode-cn.com/problems/minimum-genetic-mutation/#/description)中,将每个基因片段作为一个节点,每两个可以转换的基因,即表示了两个节点之间可达,因此可以适用于这两种搜索算法.

### 贪心算法
贪心算法适用前提是:每次选择局部最优解能够获取到整体最优解.那么在对于希望使用`贪心算法`求解的问题来说,关键是证明: *整个问题的最优解一定由在贪心策略中存在的子问题的最优解得来的*。
> 动态规划是贪心算法的泛化,贪心是动态规划的特例
动态规划本质是穷举法,但是通过记忆化缓存的方式避免对已求解的子问题重复求解,复杂度高.
贪心算法每次都对通过取当前最优解来把其他非最优解的子问题排除掉(剪枝),结果一般不是整体最优解,但复杂度低.

### 二分查找
二分查找题型中包含了`常规有序`和`旋转排序`题型
* 对于`常规有序`,直接套模版即可.
* 对于`旋转排序`,在mid的两侧,肯定是一侧有序,一侧无序.关键在于求得mid后,判断mid左侧和右侧是否有序,当知道一侧有序后,我们就可以利用有序这个新增的信息做判断,是选择有序一侧作为下一次迭代的范围还是选择无序一侧.

## 总结
本周主要学习了深度优先搜索、广度优先搜索、贪心算法和二分查找等内容.