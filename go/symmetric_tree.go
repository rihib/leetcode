//lint:file-ignore U1000 Ignore all unused code
package main

import "container/list"

func isSymmetricRecursive(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return (t1.Val == t2.Val) &&
		isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
}

func isSymmetricIterative(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := list.New()
	queue.PushBack(root.Left)
	queue.PushBack(root.Right)

	for queue.Len() > 0 {
		t1 := queue.Remove(queue.Front()).(*TreeNode)
		t2 := queue.Remove(queue.Front()).(*TreeNode)

		if t1 == nil && t2 == nil {
			continue
		}
		if t1 == nil || t2 == nil || t1.Val != t2.Val {
			return false
		}

		queue.PushBack(t1.Left)
		queue.PushBack(t2.Right)
		queue.PushBack(t1.Right)
		queue.PushBack(t2.Left)
	}

	return true
}
