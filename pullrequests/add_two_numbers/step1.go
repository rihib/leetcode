//lint:file-ignore U1000 Ignore all unused code
package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
時間：8分35秒

方針自体はすぐに決まり、繰り上げと２つの数の桁数の違いに気をつけて、足していけばいいと考えました。末端のノードの値が0になるときの対処に少し悩みました。
とりあえず動くように書いたのでコードの重複などが多く、こなれていないコードだなと思います。
*/
func addTwoNumbers_step1(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Next = new(ListNode)
	dummy.Next.Val = 0
	curr := dummy
	for l1 != nil && l2 != nil {
		curr = curr.Next
		sum := l1.Val + l2.Val + curr.Val
		curr.Val = sum % 10
		curr.Next = new(ListNode)
		curr.Next.Val = sum / 10
		l1, l2 = l1.Next, l2.Next
	}
	for l1 != nil {
		curr = curr.Next
		sum := l1.Val + curr.Val
		curr.Val = sum % 10
		curr.Next = new(ListNode)
		curr.Next.Val = sum / 10
		l1 = l1.Next
	}
	for l2 != nil {
		curr = curr.Next
		sum := l2.Val + curr.Val
		curr.Val = sum % 10
		curr.Next = new(ListNode)
		curr.Next.Val = sum / 10
		l2 = l2.Next
	}
	if curr.Next.Val == 0 {
		curr.Next = nil
	}
	return dummy.Next
}
