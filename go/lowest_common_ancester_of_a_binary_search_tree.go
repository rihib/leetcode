//lint:file-ignore U1000 Ignore all unused code
package main

func lowestCommonAncestorIterative(root, p, q *TreeNode) *TreeNode {
	node := root
	for node != nil {
		if p.Val <= node.Val && node.Val <= q.Val || q.Val <= node.Val && node.Val <= p.Val {
			return node
		}
		if p.Val < node.Val && q.Val < node.Val {
			node = node.Left
		}
		if node.Val < p.Val && node.Val < q.Val {
			node = node.Right
		}
	}
	return nil
}

func lowestCommonAncestorRecursive(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestorRecursive(root.Left, p, q)
	}
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestorRecursive(root.Right, p, q)
	}
	return root
}
