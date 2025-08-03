//lint:file-ignore U1000 Ignore all unused code
package main

// dummyを使うことで最後に値0を持つノードを作らずに済む
// 単純にl1とl2の数を求めて足し算する方法だと、64個以上のノード数を持つリスト同士の足し算ができない
func addTwoNumbersIteratice(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	nodeAdded, node1, node2 := dummy, l1, l2
	carry := 0
	for node1 != nil || node2 != nil || carry != 0 {
		val1, val2 := 0, 0
		if node1 != nil {
			val1 = node1.Val
			node1 = node1.Next
		}
		if node2 != nil {
			val2 = node2.Val
			node2 = node2.Next
		}
		sum := val1 + val2 + carry
		nodeAdded.Next = &ListNode{sum % 10, nil}
		carry = sum / 10
		nodeAdded = nodeAdded.Next
	}
	return dummy.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersWithCarry(l1, l2, 0)
}

func addTwoNumbersWithCarry(l1, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}
	sum := carry
	if l1 != nil {
		sum += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		sum += l2.Val
		l2 = l2.Next
	}
	next := addTwoNumbersWithCarry(l1, l2, sum/10)
	return &ListNode{sum % 10, next}
}
