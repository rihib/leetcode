//lint:file-ignore U1000 Ignore all unused code
package main

import "math"

type TreeBalance struct {
	isBalanced bool
	height     int
}

func isBalanced(root *TreeNode) bool {
	return checkBalance(root).isBalanced
}

func checkBalance(root *TreeNode) *TreeBalance {
	if root == nil {
		return &TreeBalance{true, 0}
	}
	left, right := checkBalance(root.Left), checkBalance(root.Right)
	isBalanced := left.isBalanced && right.isBalanced &&
		math.Abs(float64(left.height-right.height)) <= 1.0
	height := max(left.height, right.height) + 1
	return &TreeBalance{isBalanced, height}
}
