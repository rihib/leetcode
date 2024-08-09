//lint:file-ignore U1000 Ignore all unused code
package main

// `https://github.com/ryo-devz/LeetCode/pull/1#discussion_r1710718113`に書かれているとおり、印をつける方法はデメリットが大きい
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
