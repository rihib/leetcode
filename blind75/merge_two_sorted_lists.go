package blind75

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var res *ListNode
	if list1.Val <= list2.Val {
		res = list1
		list1 = list1.Next
	} else {
		res = list2
		list2 = list2.Next
	}

	curNode := res
	for {
		if list1 == nil {
			curNode.Next = list2
			return res
		}
		if list2 == nil {
			curNode.Next = list1
			return res
		}

		if list1.Val <= list2.Val {
			curNode.Next = list1
			curNode = curNode.Next
			list1 = list1.Next
		} else {
			curNode.Next = list2
			curNode = curNode.Next
			list2 = list2.Next
		}
	}
}
