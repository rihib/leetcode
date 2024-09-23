//lint:file-ignore U1000 Ignore all unused code
package lowestcommonancesterofabinarysearchtree

/*
より見やすくなるようにリファクタしました。また、再帰を使った実装もしてみました。
エラー処理についてはStep1と同様です。
*/
func lowestCommonAncestorIterative(root, p, q *TreeNode) *TreeNode {
	node := root
	for node != nil {
		if p.Val < node.Val && q.Val < node.Val {
			node = node.Left
			continue
		}
		if node.Val < p.Val && node.Val < q.Val {
			node = node.Right
			continue
		}
		return node
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
