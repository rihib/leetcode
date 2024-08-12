//lint:file-ignore U1000 Ignore all unused code
package main

func reverseList_iterative(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode
	for curr != nil {
		prev, curr, curr.Next = curr, curr.Next, prev
	}
	return prev
}

func reverseList_recurssive(head *ListNode) *ListNode {
	curr := head
	if curr == nil || curr.Next == nil {
		return curr
	}
	next := curr.Next
	reversedHead := reverseList_recurssive(next)
	curr.Next, next.Next = nil, curr
	return reversedHead
}
