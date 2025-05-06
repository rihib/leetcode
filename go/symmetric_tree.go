//lint:file-ignore U1000 Ignore all unused code
package main

import "container/list"

func isSymmetricRecursive(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

type nodePair struct {
	left  *TreeNode
	right *TreeNode
}

func isSymmetricBFS(root *TreeNode) bool {
	queue := list.New()
	queue.PushBack(nodePair{root.Left, root.Right})
	for queue.Len() > 0 {
		f := queue.Remove(queue.Front()).(nodePair)
		if f.left == nil && f.right == nil {
			continue
		}
		if f.left == nil || f.right == nil || f.left.Val != f.right.Val {
			return false
		}
		queue.PushBack(nodePair{f.left.Left, f.right.Right})
		queue.PushBack(nodePair{f.left.Right, f.right.Left})
	}
	return true
}

func isSymmetricDFS(root *TreeNode) bool {
	stack := []nodePair{{root.Left, root.Right}}
	for len(stack) > 0 {
		f := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if f.left == nil && f.right == nil {
			continue
		}
		if f.left == nil || f.right == nil || f.left.Val != f.right.Val {
			return false
		}
		stack = append(stack, nodePair{f.left.Left, f.right.Right})
		stack = append(stack, nodePair{f.left.Right, f.right.Left})
	}
	return true
}
