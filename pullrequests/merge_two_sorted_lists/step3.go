//lint:file-ignore U1000 Ignore all unused code
package mergetwosortedlists

/*
	感想：
		現在（6月）解いたコードになります。大分余裕を持って書けるようになってきました。一番初めはポ
		インタの向きがよくわからなくなるということがあったのですが、そういったこともなくなり、ポイ
		ンタの操作を余裕を持って書けるようになってきました。暗記ではなく、自然と同じようなコードが
		出てくるようになりました。またGo自体も初期に比べて無理なく自然に書けるようになりました。

		Step2に比べて変数名が改善されています（curではなくtailに変更）。また個人的にシンプルで見
		やすいと思っているので多重代入をよく使うようになりました。

		また、個人的な感覚として、list1またはlist2のどちらか片方がnilの場合の処理をfor文の中に入
		れるか、外に出すかをその時の気分で意図的に変えるようになりました。

	質問したいこと：
		- tailという変数名についてどう思いますか？
		- 多重代入は使った方がシンプルで見やすいと個人的に思ってるのですがどう思いますか？もちろん
			互いに影響を与え合ってしまうような変数同士の場合は挙動がundefinedだと思うので避けるべき
			だとは思いますが。
		- 個人的には上と下の回答はどちらも大差ないと思っていてどちらでも良いと思っているのですが、
			どう思いますか？
*/

func mergeTwoLists_step3(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := new(ListNode)
	tail := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next, list1 = list1, list1.Next
		} else {
			tail.Next, list2 = list2, list2.Next
		}
		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	}
	if list2 != nil {
		tail.Next = list2
	}

	return dummy.Next
}

func mergeTwoLists_step3_anothersolution(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := new(ListNode)
	tail := dummy

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

	return dummy.Next
}
