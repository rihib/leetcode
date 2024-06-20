//lint:file-ignore U1000 Ignore all unused code
package main

import "math"

type TreeBalance struct {
	IsBalanced bool
	Height     int
}

func isBalanced(root *TreeNode) bool {
	return checkBalance(root).IsBalanced
}

func checkBalance(root *TreeNode) TreeBalance {
	if root == nil {
		return TreeBalance{true, 0}
	}

	l, r := checkBalance(root.Left), checkBalance(root.Right)
	b := l.IsBalanced && r.IsBalanced &&
		math.Abs(float64(l.Height-r.Height)) <= 1.0
	h := max(l.Height, r.Height) + 1
	return TreeBalance{b, h}
}
