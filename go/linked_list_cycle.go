//lint:file-ignore U1000 Ignore all unused code
package main

func hasCycleFloyd(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func hasCycleMap(head *ListNode) bool {
	seen := make(map[*ListNode]struct{})
	node := head
	for node != nil {
		if _, ok := seen[node]; ok {
			return true
		}
		seen[node] = struct{}{}
		node = node.Next
	}
	return false
}
