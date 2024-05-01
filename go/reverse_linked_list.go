//lint:file-ignore U1000 Ignore all unused code
package main

func reverseList(head *ListNode) *ListNode {
	// Test Case: [], [1], [1, 2], [1, 2, 3], [1, 2, 3, 4]
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}

	prev := head
	next := head.Next
	head.Next = nil

	for next != nil {
		nextNext := next.Next
		next.Next = prev
		prev = next
		next = nextNext
	}
	return prev
}
