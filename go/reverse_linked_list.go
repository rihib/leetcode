//lint:file-ignore U1000 Ignore all unused code
package main

func reverseListIterative(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev, curr = curr, next
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
