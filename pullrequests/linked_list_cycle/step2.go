//lint:file-ignore U1000 Ignore all unused code
package linkedlistcycle

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
他に既に訪ねたノードのリストを保持しておく方法も考えられる。
ノードのポインタ（メモリアドレス）をキーとしている。
*/
func hasCycleMap(head *ListNode) bool {
	seen := make(map[*ListNode]struct{})
	node := head
	for node != nil {
		if _, ok := seen[node]; ok {
			return true
		}
		seen[node] = struct{}{}
		node = node.Next
	}
	return false
}
