//lint:file-ignore U1000 Ignore all unused code
package mergetwosortedlists

/*
	個人的にLeetCodeを進めていて、2月ごろに解いたもの。GitHubのコミット履歴から引っ張ってきま
	した。

	解いた時間を記録していなかったのですが、恐らく20~30分ぐらいかかった気がします。解いている途
	中でポインタの向きがよくわからなくなってしまい、時間がかかってしまった記憶があります。また、
	Goの勉強も兼ねて書いたので、言語自体の知識も少ない状態で書きました。

	綺麗に書いている余裕がなかったので、とりあえず動くコードを書いたという感じです。とりあえず値
	の小さいノードから順番にポインタの向きを変えてマージしていけば良いとはわかったのですが、実装
	力に余裕がなく、コードが冗長になってしまったという印象です。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists_step1(list1 *ListNode, list2 *ListNode) *ListNode {
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
