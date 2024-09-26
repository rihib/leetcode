//lint:file-ignore U1000 Ignore all unused code
package main

func middleNode(head *ListNode) *ListNode {
	count := 0
	node := head
	for node != nil {
		node = node.Next
		count++
	}
	middle := head
	for i := 0; i < count/2; i++ {
		middle = middle.Next
	}
	return middle
}

func middleNode2(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}
