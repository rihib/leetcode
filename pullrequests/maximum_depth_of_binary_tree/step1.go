//lint:file-ignore U1000 Ignore all unused code
package maximumdepthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
時間：2分

一番オーソドックスなやり方。

１フレームの大きさは、引数8 + 返り値8 + ローカル変数(8 + 8) + FP + 戻りアドレス = 48Bぐらい。最大ノード数が10^4なのでスタックサイズは、48 * 10^4 = 480KBぐらいになると予想できる。
Linuxのデフォルトのスタックの大きさが8MBなのでスタックオーバーフローの危険性は低い。
goroutineのデフォルトのスタックの大きさは8KBだが、動的に拡張されるようになっているので通常、スタックサイズがGBレベルになっても問題なく再帰を回すことができる。
https://github.com/rihib/leetcode/pull/15#issue-2459505166
*/
func maxDepthRecursive(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepthRecursive(root.Left)
	rightDepth := maxDepthRecursive(root.Right)
	return max(leftDepth, rightDepth) + 1
}
