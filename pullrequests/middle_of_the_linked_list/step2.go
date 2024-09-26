//lint:file-ignore U1000 Ignore all unused code
package middleofthelinkedlist

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
最初の早期リターンはなくても問題ないことに気づき、削除した。
*/
func middleNodeStep2(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}
