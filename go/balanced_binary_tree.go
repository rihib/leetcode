//lint:file-ignore U1000 Ignore all unused code
package main

import "math"

type TreeBalance struct {
	IsBalanced bool
	Height     int
}

func isBalanced(root *TreeNode) bool {
	return dfs(root).IsBalanced
}

func dfs(root *TreeNode) TreeBalance {
	if root == nil {
		return TreeBalance{true, 0}
	}

	l, r := dfs(root.Left), dfs(root.Right)
	b := (l.IsBalanced &&
		r.IsBalanced &&
		(int(math.Abs(float64(l.Height)-float64(r.Height))) <= 1))
	h := max(l.Height, r.Height) + 1
	return TreeBalance{b, h}
}
