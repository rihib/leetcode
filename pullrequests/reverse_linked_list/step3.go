//lint:file-ignore U1000 Ignore all unused code
package reverselinkedlist

/*
	時間: 2分、2分
	イテレーティブな方法と再帰を使った方法の２通りで解きました。

	質問：
		- 20行目について、１行にまとめているのですが、読みやすさなどの点ではどう思いますか？

		- 与えられたhead自体を操作するのは避けた方が良いか？（例えばcurrなどの別の変数に代入して
			から操作した方が良いか）

		- 変数名としてreversedHeadか、reversedListHeadのどちらの方が良いか？
			- reversedListHeadだとこれぐらいのコードの短さで使うにしては少し長すぎるかなと思った
				のでここではreversedHeadでも十分通じると思い、これにしました

	他の人のPRを見て思ったこと
		- スタックを使って解く方法も確かにあるなと思いました。パフォーマンス的には再帰を使って解く
			方法とほぼ同じ？（再帰もスタックを使っていると言えば使っている）
*/

func reverseList_iterative_step3(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		prev, head, head.Next = head, head.Next, prev
	}
	return prev
}

func reverseList_recurssive_step3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	reversedHead := reverseList_recurssive_step3(next)
	head.Next, next.Next = nil, head
	return reversedHead
}
