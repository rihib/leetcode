//lint:file-ignore U1000 Ignore all unused code
package main

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	turtle, rabbit := head.Next, head.Next.Next
	for turtle != rabbit {
		if rabbit == nil || rabbit.Next == nil {
			return false
		}

		turtle = turtle.Next
		rabbit = rabbit.Next.Next
	}
	return true
}
