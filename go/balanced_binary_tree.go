//lint:file-ignore U1000 Ignore all unused code
package main

import "math"

type TreeBalance struct {
	IsBalanced bool
	Height     int
}

func isBalanced(root *TreeNode) bool {
	return traversal(root).IsBalanced
}

func traversal(root *TreeNode) TreeBalance {
	if root == nil {
		return TreeBalance{true, 0}
	}

	l, r := traversal(root.Left), traversal(root.Right)
	b := l.IsBalanced &&
		r.IsBalanced &&
		(int(math.Abs(float64(l.Height)-float64(r.Height))) <= 1)
	h := max(l.Height, r.Height) + 1
	return TreeBalance{b, h}
}
