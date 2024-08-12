//lint:file-ignore U1000 Ignore all unused code
package reverselinkedlist

func reverseList_iterative_step4(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode
	for curr != nil {
		prev, curr, curr.Next = curr, curr.Next, prev
	}
	return prev
}

func reverseList_recurssive_step4(head *ListNode) *ListNode {
	curr := head
	if curr == nil || curr.Next == nil {
		return curr
	}
	next := curr.Next
	reversedHead := reverseList_recurssive_step4(next)
	curr.Next, next.Next = nil, curr
	return reversedHead
}
