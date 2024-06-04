//lint:file-ignore U1000 Ignore all unused code
package main

func maxDepth(root *TreeNode) int {
	// TestCase: [], [0], [0, 1], [0, 1, 2], [0, 1, null, 2]
	if root == nil {
		return 0
	}

	maxdepth := max(maxDepth(root.Left), maxDepth(root.Right)) + 1
	return maxdepth
}
