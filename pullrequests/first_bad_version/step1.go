//lint:file-ignore U1000 Ignore all unused code
package firstbadversion

const BadVersion = 4

func isBadVersion(version int) bool {
	return version == BadVersion
}

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
時間：8分30秒

単なる二分探索だが、少し混乱してしまったため、時間がかかってしまった。二分探索については理解しているつもりだが、考えずにできるレベルにはまだなれてないので少し問題文が変更されていると頭で考える必要が出てきてしまう。

半閉区画の場合、rightは常にtrueの範囲に収まり続け、leftはfalseとtrueの境界にあるとき、mid+1によりtrueに移動するため、left == rightになった瞬間はleftはFirst bad versionであることがわかる。
leftがmid+1でtrueに移動したときにrightがleftよりも大きい位置（left < right）にいたとしても、探索は続くが、leftもrightもtrueであればmidも必ずtrueになるため、leftは動かず、rightだけが徐々にleftに近づいていき、最終的にはleftと合流する。
right = midと更新される以上、rightがleftを飛び越えることはなくtrueの範囲にとどまり続け、left == rightになるまで探索は終了しないので、firstBadVersionHalfClosed2のようにleftの代わりにrightを返しても問題ない。
rightがleftを飛び越えることはないのでfirstBadVersionHalfClosed3のように書くこともできる。

閉区画の場合、falseとtrueの境界にあるとき、rightはmid-1によりfalseの方に移動するのに対し、leftはtrueの方に移動するため、right < leftになった瞬間にleftはFirst bad versionであることがわかる。left == rightの場合、trueの範囲内でleft == rightになる場合は問題ないが、falseの範囲内でleft == rightとなる場合は探索結果になってしまうため、right < leftになるまで探索を続ける必要がある。
right < leftになった瞬間に探索が終了するわけだが、このとき、leftはtrue、rightはfalseになっているため、閉区画の場合は必ずleftを返す必要がある。
*/
func firstBadVersionHalfClosed1(n int) int {
	left, right := 1, n+1
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func firstBadVersionHalfClosed2(n int) int {
	left, right := 1, n+1
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func firstBadVersionHalfClosed3(n int) int {
	left, right := 1, n+1
	for {
		if left == right {
			break
		}
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func firstBadVersionClosed(n int) int {
	left, right := 1, n
	for left <= right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
