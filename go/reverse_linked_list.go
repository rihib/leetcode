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
	if head == nil || head.Next == nil {
		return head
	}
	curr, next := head, head.Next
	reversed := reverseListRecurssive(next)
	next.Next = curr
	curr.Next = nil
	return reversed
}
