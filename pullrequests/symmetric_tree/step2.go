//lint:file-ignore U1000 Ignore all unused code
package symmetrictree

import "container/list"

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
イテレーティブにも解いてみた。BFSなのでキューを使って実装した。
*/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := list.New()
	queue.PushBack(root.Left)
	queue.PushBack(root.Right)
	for queue.Len() > 0 {
		left := queue.Remove(queue.Front()).(*TreeNode)
		right := queue.Remove(queue.Front()).(*TreeNode)
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		queue.PushBack(left.Left)
		queue.PushBack(right.Right)
		queue.PushBack(left.Right)
		queue.PushBack(right.Left)
	}
	return true
}
