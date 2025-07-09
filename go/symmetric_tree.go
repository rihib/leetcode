//lint:file-ignore U1000 Ignore all unused code
package main

import "container/list"

func isSymmetric(root *TreeNode) bool {
	return areSymmetric(root.Left, root.Right)
}

func areSymmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	return areSymmetric(left.Left, right.Right) && areSymmetric(left.Right, right.Left)
}

type nodePair struct {
	left  *TreeNode
	right *TreeNode
}

func isSymmetricBFS(root *TreeNode) bool {
	queue := list.New()
	queue.PushBack(nodePair{root.Left, root.Right})
	for queue.Len() > 0 {
		pair := queue.Remove(queue.Front()).(nodePair)
		if pair.left == nil && pair.right == nil {
			continue
		}
		if pair.left == nil || pair.right == nil || pair.left.Val != pair.right.Val {
			return false
		}
		queue.PushBack(nodePair{pair.left.Left, pair.right.Right})
		queue.PushBack(nodePair{pair.left.Right, pair.right.Left})
	}
	return true
}

func isSymmetricDFS(root *TreeNode) bool {
	stack := []nodePair{{root.Left, root.Right}}
	for len(stack) > 0 {
		pair := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pair.left == nil && pair.right == nil {
			continue
		}
		if pair.left == nil || pair.right == nil || pair.left.Val != pair.right.Val {
			return false
		}
		stack = append(stack, []nodePair{
			{pair.left.Left, pair.right.Right},
			{pair.left.Right, pair.right.Left},
		}...)
	}
	return true
}
