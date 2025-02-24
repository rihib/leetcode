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

func isSymmetricIterative(root *TreeNode) bool {
	queue := list.New()
	pair := nodePair{root, root}
	queue.PushBack(pair)
	for queue.Len() > 0 {
		pair := queue.Remove(queue.Front()).(nodePair)
		left, right := pair.left, pair.right
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		queue.PushBack(nodePair{left.Left, right.Right})
		queue.PushBack(nodePair{left.Right, right.Left})
	}
	return true
}
