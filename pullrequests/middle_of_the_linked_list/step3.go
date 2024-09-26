//lint:file-ignore U1000 Ignore all unused code
package middleofthelinkedlist

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
他の方の解法を見て実装してみた。この方法を先に思いつかなかったのが逆に良くない気がした。
*/
func middleNodeStep3(head *ListNode) *ListNode {
	count := 0
	node := head
	for node != nil {
		node = node.Next
		count++
	}
	middle := head
	for i := 0; i < count/2; i++ {
		middle = middle.Next
	}
	return middle
}
