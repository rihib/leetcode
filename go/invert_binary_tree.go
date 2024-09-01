//lint:file-ignore U1000 Ignore all unused code
package main

func invertTreeRecursive(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = invertTreeRecursive(root.Right), invertTreeRecursive(root.Left)
	}
	return root
}

func invertTreeIterative(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return root
}
