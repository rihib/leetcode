//lint:file-ignore U1000 Ignore all unused code
package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := new(ListNode)
	cur := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next, list1 = list1, list1.Next
		} else {
			cur.Next, list2 = list2, list2.Next
		}
		cur = cur.Next
	}

	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}

	return dummy.Next
}

func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	sorted := new(ListNode)
	tail := sorted

	for list1 != nil || list2 != nil {
		if list1 == nil {
			tail.Next = list2
			break
		}
		if list2 == nil {
			tail.Next = list1
			break
		}

		if list1.Val < list2.Val {
			tail.Next, list1 = list1, list1.Next
		} else {
			tail.Next, list2 = list2, list2.Next
		}
		tail = tail.Next
		tail.Next = nil
	}

	return sorted.Next
}
