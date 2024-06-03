//lint:file-ignore U1000 Ignore all unused code
package main

func middleNode(head *ListNode) *ListNode {
	// TestCase: [], [1], [1, 2], [1, 2, 3], [1, 2, 3, 4]
	if head == nil {
		return nil
	}

	l := getListLen(head)
	mid_idx := l / 2
	mid := head
	for i := 0; i < mid_idx; i++ {
		mid = mid.Next
	}
	return mid
}

func getListLen(head *ListNode) int {
	if head == nil {
		return 0
	}

	l := 0
	for head != nil {
		l++
		head = head.Next
	}

	return l
}
