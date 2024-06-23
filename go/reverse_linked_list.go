//lint:file-ignore U1000 Ignore all unused code
package main

func reverseList_iterative(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		prev, head, head.Next = head, head.Next, prev
	}
	return prev
}

func reverseList_recurssive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	reversedHead := reverseList(next)
	head.Next, next.Next = nil, head
	return reversedHead
}
