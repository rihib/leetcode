//lint:file-ignore U1000 Ignore all unused code
package main

func maxDepthRecursive(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepthRecursive(root.Left)
	rightDepth := maxDepthRecursive(root.Right)
	return max(leftDepth, rightDepth) + 1
}

func maxDepthIterativeDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maximum := 0
	stack := []entry{{root, 1}}
	for len(stack) > 0 {
		e := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		maximum = max(maximum, e.depth)
		if e.node.Left != nil {
			stack = append(stack, entry{e.node.Left, e.depth + 1})
		}
		if e.node.Right != nil {
			stack = append(stack, entry{e.node.Right, e.depth + 1})
		}
	}
	return maximum
}

func maxDepthIterativeBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maximum := 0
	queue := []entry{{root, 1}}
	for len(queue) > 0 {
		e := queue[0]
		queue = queue[1:]
		maximum = max(maximum, e.depth)
		if e.node.Left != nil {
			queue = append(queue, entry{e.node.Left, e.depth + 1})
		}
		if e.node.Right != nil {
			queue = append(queue, entry{e.node.Right, e.depth + 1})
		}
	}
	return maximum
}

type entry struct {
	node  *TreeNode
	depth int
}
