//lint:file-ignore U1000 Ignore all unused code
package main

// dummyを使うことで最後に値0を持つノードを作らずに済む
// 単純にl1とl2の数を求めて足し算する方法だと、64個以上のノード数を持つリスト同士の足し算ができない
func addTwoNumbersIteratice(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	current, node1, node2 := dummy, l1, l2
	carry := 0
	for node1 != nil || node2 != nil || carry != 0 {
		sum := carry
		if node1 != nil {
			sum += node1.Val
			node1 = node1.Next
		}
		if node2 != nil {
			sum += node2.Val
			node2 = node2.Next
		}
		current.Next = &ListNode{sum % 10, nil}
		current = current.Next
		carry = sum / 10
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
	sum := carry
	if l1 != nil {
		sum += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		sum += l2.Val
		l2 = l2.Next
	}
	next := addTwoNumbersHelper(l1, l2, sum/10)
	return &ListNode{sum % 10, next}
}
