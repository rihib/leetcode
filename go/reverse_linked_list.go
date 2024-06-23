//lint:file-ignore U1000 Ignore all unused code
package main

func reverseList_iterative(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		next := head.Next
		head.Next = prev
		prev, head = head, next
	}

	return prev
}

func reverseList_recursive(head *ListNode) *ListNode {
	// Test Case: [], [1], [1, 2], [1, 2, 3], [1, 2, 3, 4]
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	reversedListHead := reverseList_recursive(next)
	next.Next, head.Next = head, nil
	return reversedListHead
}
