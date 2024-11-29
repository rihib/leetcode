//lint:file-ignore U1000 Ignore all unused code
package addtwonumbers

/*
carryを導入した方がより明確でシンプルに書けると気づき変更しました。
carryを導入したことにより、末端に余計なノード（値が0のノード）が発生しなくなりました。
*/
func addTwoNumbers_step3(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	curr := dummy
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		curr.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		curr = curr.Next
	}
	return dummy.Next
}
