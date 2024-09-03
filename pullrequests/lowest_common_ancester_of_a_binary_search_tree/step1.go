//lint:file-ignore U1000 Ignore all unused code
package lowestcommonancesterofabinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
時間：6分30秒

ルートから見ていく時、共通の祖先になっていないときは必ず、pとqは左か右のどちらかの同じ部分木にいるはず。pとqの間の値になったら共通の祖先になったとわかる。

本来は見つからなかった場合は返り値としてerrorを返したかったのですが、LeetCodeの制約上変えられないのでnilを返すようにしています。
*/
func lowestCommonAncestorIterativeStep1(root, p, q *TreeNode) *TreeNode {
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
