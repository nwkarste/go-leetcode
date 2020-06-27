package utils

import ()

//TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var stack []*StackNode
	stack = append(stack, &StackNode{Node: root})
	var traverse []int
	count := 0
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		if node.Node == nil {
			stack = stack[:len(stack)-1]
			traverse = append(traverse, 1)
			if count > 4 {
				return traverse
			}
		} else if node.Seen == 1 {
			stack = stack[:len(stack)-1]
			traverse = append(traverse, node.Node.Val)
			stack = append(stack, &StackNode{Node: node.Node.Right})
			traverse = append(traverse, 2)
			if count > 4 {
				return traverse
			}
			return traverse
		} else {
			traverse = append(traverse, node.Seen)
			node.Seen = 1
			traverse = append(traverse, node.Seen)
			stack = append(stack, &StackNode{Node: node.Node.Left})
			traverse = append(traverse, 3)
			if count > 4 {
				return traverse
			}
		}
		count++
	}
	return traverse
}

type StackNode struct {
	Seen int
	Node *TreeNode
}
