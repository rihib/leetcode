//lint:file-ignore U1000 Ignore all unused code
package mergetwosortedlists

/*
	他の人のコードを色々見て、冗長な部分を消し、リファクタをしました。
*/

func mergeTwoLists_step2(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := new(ListNode)
	cur := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}

	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}

	return dummy.Next
}
