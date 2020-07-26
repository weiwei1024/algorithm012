package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//236.二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pList := make([]*TreeNode, 0) //p节点父节点列表
	qList := make([]*TreeNode, 0) //q节点父节点列表
	searchParent(root, p, &pList)
	searchParent(root, q, &qList)

	//两条链表是否有相交节点
	i, j := len(pList)-1, len(qList)-1
	for i >= 0 && j >= 0 {
		if pList[i] != qList[i-1] {
			break
		} else {
			i--
			j--
		}
	}
	if i < 0 {
		return pList[i+1]
	}
	if j < 0 {
		return qList[j+1]
	}
	return root
}

//查一个节点的父亲节点列表
func searchParent(root, n *TreeNode, parentList *[]*TreeNode) bool {
	if root == nil {
		return false
	}
	if root == n {
		*parentList = append(*parentList, n)
		return true
	}

	lt := searchParent(root.Left, n, parentList)
	rt := searchParent(root.Right, n, parentList)
	if lt || rt {
		*parentList = append(*parentList, root)
	}
	return lt || rt
}
