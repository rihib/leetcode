//lint:file-ignore U1000 Ignore all unused code
package main

func reverseListIterative(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode
	for curr != nil {
		prev, curr, curr.Next = curr, curr.Next, prev
	}
	return prev
}

func reverseListRecurssive(head *ListNode) *ListNode {
	curr := head
	if curr == nil || curr.Next == nil {
		return curr
	}
	next := curr.Next
	reversedHead := reverseListRecurssive(next)
	curr.Next, next.Next = nil, curr
	return reversedHead
}
