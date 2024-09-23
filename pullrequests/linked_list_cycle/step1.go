//lint:file-ignore U1000 Ignore all unused code
package linkedlistcycle

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
時間：5分

この問題を見たことがあり、フロイドの循環検出法を知っていたため、その方法で解いてみた。
`slow == fast`ではポインタ同士を比較しているため、同一のオブジェクトであることを確認できる。
*/
func hasCycleFloyd(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
