//lint:file-ignore U1000 Ignore all unused code
package main

func reverseList1(head *ListNode) *ListNode {
	// Test Case: [], [1], [1, 2], [1, 2, 3], [1, 2, 3, 4]
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
