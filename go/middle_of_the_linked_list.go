//lint:file-ignore U1000 Ignore all unused code
package main

func middleNode(head *ListNode) *ListNode {
	length := 0
	node := head
	for node != nil {
		length++
		node = node.Next
	}
	node = head
	for i := 0; i < length/2; i++ {
		node = node.Next
	}
	return node
}

func middleNode2(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}
