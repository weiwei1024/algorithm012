package tree

//105.构建二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	return buildTreeHelp(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

//辅助函数
func buildTreeHelp(preorder []int, inorder []int, pSt, pEd, iSt, iEd int) *TreeNode {
	if pSt > pEd {
		return nil
	}
	pRoot := preorder[pSt]
	//定位中序遍历根节点(可优化)
	index := 0
	for i := iSt; i <= iEd; i++ {
		if pRoot == inorder[i] {
			index = i
		}
	}
	leftTreeSize := index - iSt
	lNode := buildTreeHelp(preorder, inorder, pSt+1, pSt+leftTreeSize, iSt, index-1) //构建左子树
	rNode := buildTreeHelp(preorder, inorder, pSt+leftTreeSize+1, pEd, index+1, iEd) //构建右子树
	//返回当前节点
	return &TreeNode{
		Val:   inorder[index],
		Left:  lNode,
		Right: rNode,
	}
}
