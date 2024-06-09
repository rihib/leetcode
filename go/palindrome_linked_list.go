//lint:file-ignore U1000 Ignore all unused code
package main

func isPalindromeLinkedList(head *ListNode) bool {
	// TestCase: [], [1], [1, 1], [1, 2, 1], [1, 2, 2, 1]
	if head == nil {
		return true
	}

	slow, fast := head, head
	var val_list []int

	for fast != nil && fast.Next != nil {
		val_list = append(val_list, slow.Val)
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	for i := len(val_list) - 1; i >= 0; i-- {
		if slow == nil || val_list[i] != slow.Val {
			return false
		}
		slow = slow.Next
	}

	return true
}
