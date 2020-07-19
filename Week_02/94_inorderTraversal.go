package Week_02

import "errors"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//二叉树的中序非递归遍历
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	s := &Stack{}
	result := []int{}
	if root == nil {
		return nil
	}
	curr := root
	for curr != nil || s.Len() > 0 {
		//左节点不为空,将左节点入栈
		for curr != nil {
			_ = s.Push(curr)
			curr = curr.Left
		}
		//左节点为空,出栈
		v, err := s.Pop()
		if err != nil {
			break
		}
		curr = v.(*TreeNode)
		result = append(result, curr.Val)
		curr = curr.Right
	}
	return result
}

type Stack []interface{}

//入栈
func (stack *Stack) Push(v interface{}) error {
	*stack = append(*stack, v)
	return nil
}

//出栈
func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack
	if len(theStack) <= 0 {
		return nil, errors.New("stack empty")
	}
	value := theStack[len(*stack)-1]
	*stack = theStack[:len(theStack)-1]
	return value, nil
}

//返回栈高度
func (stack *Stack) Len() int {
	return len(*stack)
}
