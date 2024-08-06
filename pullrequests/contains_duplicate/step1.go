//lint:file-ignore U1000 Ignore all unused code
package containsduplicate

/*
かなり前に解いたときのもの（コミット履歴からとってきました）なので
詳細は忘れてしまいました（すいません）。
基本的にはマップを使って重複を見つけようとしています。

- Step1の要素数が1以下の場合の処理はなくても変わらないことに気づいたのでStep2ではなくしています。
- valueの意味でvを使ったのですが、nのほうがnumsと対応していてより良いかなと思ったので同様にStep2ではnを使うようにしました。
- nmもnumsMapのつもりで使ったのですが同様に他にマップが使われているわけではないのでStep2では単にmとしました。
*/
func containsDuplicate_step1(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	nm := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := nm[v]; ok {
			return true
		}
		nm[v] = struct{}{}
	}
	return false
}
