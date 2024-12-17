//lint:file-ignore U1000 Ignore all unused code
package palindromelinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
時間：５分
スライスに変換し、Palindromeかどうかを比較するようにした。
*/
func isPalindromeLinkedList(head *ListNode) bool {
	var values []int
	node := head
	for node != nil {
		values = append(values, node.Val)
		node = node.Next
	}
	left, right := 0, len(values)-1
	for left < right {
		if values[left] != values[right] {
			return false
		}
		left++
		right--
	}
	return true
}
