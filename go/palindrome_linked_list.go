//lint:file-ignore U1000 Ignore all unused code
package main

func isPalindromeLinkedList(head *ListNode) bool {
	var values []int
	node := head
	for node != nil {
		values = append(values, node.Val)
		node = node.Next
	}
	left, right := 0, len(values)-1
	for left < right {
		if values[left] != values[right] {
			return false
		}
		left++
		right--
	}
	return true
}
