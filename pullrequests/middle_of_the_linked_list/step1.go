//lint:file-ignore U1000 Ignore all unused code
package middleofthelinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
時間：3分

しばらく解法が思いつかなかったが、フロイドの循環検出法を思いだし、それと似た方法で解くことができると思いついた。
*/
func middleNodeStep1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}
