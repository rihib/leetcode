//lint:file-ignore U1000 Ignore all unused code
package firstbadversion

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
実はこの問題はfirstBadVersion1のようにrightをn+1で初期化しなくてもパスすることができる。

通常の二分探索では半閉区画で書く場合、rightは常に探索済みの範囲に入れておく必要がある。もしもrightをlen(nums)-1で初期化してしまうと、一番最後の要素が探しているtargetの場合に、left < rightによってleftがtargetに到達した瞬間に探索が終了し、targetが存在しなかったという扱いになってしまう。

しかし、今回の場合は必ずFirst Bad Versionが存在する前提になっているため、left < rightによって探索が終了したらleft（またはright）が返るようになっているので、rightをnで初期化してもパスする。

しかし、本来であればBad versionが１つも存在しない入力が来ることも考えられるため、その場合の処理も加える必要がある。仮に見つからなかった場合に-1を返すようにする（もっとちゃんとするならerrorを返した方がわかりやすいとは思うが）なら、firstBadVersion2またはfirstBadVersion3のようにする必要がある。通常の二分探索に寄せて書くのであれば、firstBadVersion4のように書くことができる。

逆に言えば通常の二分探索もfirstBadVersion2のように書くことができる。この場合は単純にtarget以上か否かでfalse, trueの範囲を分け、その境界の最初のtrueを求めるようにし、最後にその求めた値がtargetかどうかを確かめれば良い。

```Go

	func binarySearch(nums []int, target int) int {
		left, right := 0, len(nums)
		for left < right {
			mid := left + (right-left)/2
			if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid
			}
		}
		if left >= len(nums) || nums[left] != target {
			return -1
		}
		return left
	}

```
*/
func firstBadVersion1(n int) int {
	left, right := 1, n
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

func firstBadVersion2(n int) int {
	left, right := 1, n+1
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if left > n || !isBadVersion(left) {
		return -1
	}
	return left
}

func firstBadVersion3(n int) int {
	left, right := 1, n+1
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if right > n || !isBadVersion(right) {
		return -1
	}
	return right
}

func firstBadVersion4(n int) int {
	left, right := 1, n+1
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) && (mid == 1 || !isBadVersion(mid-1)) {
			return mid
		}
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return -1
}
