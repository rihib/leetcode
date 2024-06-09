//lint:file-ignore U1000 Ignore all unused code
package main

func isPalindromeLinkedList(head *ListNode) bool {
	slow, fast := head, head
	var rev *ListNode

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		tmp := slow.Next
		slow.Next = rev
		rev = slow
		slow = tmp
	}

	if fast != nil {
		slow = slow.Next
	}

	for slow != nil {
		if slow.Val != rev.Val {
			return false
		}
		slow, rev = slow.Next, rev.Next
	}

	return true
}
