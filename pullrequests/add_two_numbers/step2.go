//lint:file-ignore U1000 Ignore all unused code
package addtwonumbers

/*
コードの重複を排除しました。ただまだ無駄な処理が多いです。
*/
func addTwoNumbers_step2(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Next = new(ListNode)
	dummy.Next.Val = 0
	curr := dummy
	for l1 != nil || l2 != nil {
		curr = curr.Next
		sum := curr.Val
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		curr.Val = sum % 10
		curr.Next = new(ListNode)
		curr.Next.Val = sum / 10
	}
	if curr.Next.Val == 0 {
		curr.Next = nil
	}
	return dummy.Next
}
