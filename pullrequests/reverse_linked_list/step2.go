//lint:file-ignore U1000 Ignore all unused code
package reverselinkedlist

/*
COMMENT
*/

func reverseList_step2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}

	prev := head
	curr := head.Next
	head.Next = nil

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
