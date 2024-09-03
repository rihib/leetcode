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

該当ノードが見つからなかったからといってプログラムの実行の継続が困難になるわけではないと思うので、panicやlog.Panicを使うのはやり過ぎだと思います。
場合によっては絶対に見つからないことが起こる入力はしないはずだと言える状況であればlog.Fatalを使ってログに書き込んだ後にos.Exit(1)を呼び出してプログラムを終了させるのが良い可能性もありますが、それも通常の場合やりすぎな気がします。
他にはlog.Printなどを使ってログに見つからなかったことを記録しても良いかもしれませんが、見つからないことが起こる入力があり得るのであれば単にerrorを返り値として返して呼び出し側で処理するのが良いと思っています。
他には単に見つからなかった場合はnilを返すなども手としてはあると思いますが、そうする場合は呼び出し側にもわかるようにコメント等で書いておいて欲しいなと思います。
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
