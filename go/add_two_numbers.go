//lint:file-ignore U1000 Ignore all unused code
package main

// dummyを使うことで最後に値0を持つノードを作らずに済む
// 単純にl1とl2の数を求めて足し算する方法だと、64個以上のノード数を持つリスト同士の足し算ができない
func addTwoNumbersIteratice(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	node := dummy
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
		node.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		node = node.Next
	}
	return dummy.Next
}

func addTwoNumbersRecursive(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersHelper(l1, l2, 0)
}

func addTwoNumbersHelper(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}
	l1Val, l2Val := 0, 0
	if l1 != nil {
		l1Val = l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		l2Val = l2.Val
		l2 = l2.Next
	}
	sum := l1Val + l2Val + carry
	head := &ListNode{Val: sum % 10, Next: addTwoNumbersHelper(l1, l2, sum/10)}
	return head
}
