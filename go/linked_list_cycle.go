//lint:file-ignore U1000 Ignore all unused code
package main

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
