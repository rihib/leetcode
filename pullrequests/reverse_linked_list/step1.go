//lint:file-ignore U1000 Ignore all unused code
package reverselinkedlist

/*
COMMENT
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList_step1(head *ListNode) *ListNode {
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
