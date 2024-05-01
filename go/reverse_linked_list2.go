//lint:file-ignore U1000 Ignore all unused code
package main

func reverseList2(head *ListNode) *ListNode {
	// Test Case: [], [1], [1, 2], [1, 2, 3], [1, 2, 3, 4]
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	reversedListHead := reverseList2(next)
	next.Next, head.Next = head, nil
	return reversedListHead
}
