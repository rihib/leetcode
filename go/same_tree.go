//lint:file-ignore U1000 Ignore all unused code
package main

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// TestCase: [[], []], [[1], []], [[1], [1]], [[1, 2, 3], [1, 2]]
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
