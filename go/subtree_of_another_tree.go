//lint:file-ignore U1000 Ignore all unused code
package main

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// TestCase: [[], []], [[1], []], [[], [1]], [[1], [1]], [[1], [1, 2]],
	// [[1, 2, 2, nil, nil, 3], [2, 3]]
	if root == nil {
		return subRoot == nil
	}

	return areEqualTrees(root, subRoot) ||
		isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func areEqualTrees(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil || t1.Val != t2.Val {
		return false
	}
	return areEqualTrees(t1.Left, t2.Left) &&
		areEqualTrees(t1.Right, t2.Right)
}
