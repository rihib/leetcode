//lint:file-ignore U1000 Ignore all unused code
package main

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// TestCase: [[], []], [[1], []], [[], [1]], [[1], [1]], [[1], [1, 2]],
	// [[1, 2, 2, nil, nil, 3], [2, 3]]
	if root == nil {
		return subRoot == nil
	}

	return isPartOfSubtree(root, subRoot) ||
		isSubtree(root.Left, subRoot) ||
		isSubtree(root.Right, subRoot)
}

func isPartOfSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if root == nil || subRoot == nil || root.Val != subRoot.Val {
		return false
	}
	return isPartOfSubtree(root.Left, subRoot.Left) &&
		isPartOfSubtree(root.Right, subRoot.Right)
}
