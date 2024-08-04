//lint:file-ignore U1000 Ignore all unused code
package containsduplicate

/*
時間：１分
よりシンプルに書くようにしました。
また他の解法として最後に要素数を比較する方法もあると思います。

質問
- m, nという変数名をどう思いますか？（個人的にはこれぐらいの短いコードであれば逆に長い変数名をつけるのも良くないのではと思ったためあえて一文字の変数を使いました）
- ２つの解法のうちどちらのがほうが良いと思いますか？（個人的には下の解法はコード数が短くなるものの、必ずnumsを全て見ないといけないので、その点で若干パフォーマンスは劣るのかなと思っています。気にしなくても良いぐらいの話かもしれませんが、、）
*/
func containsDuplicate_step2(nums []int) bool {
	m := make(map[int]struct{})
	for _, n := range nums {
		if _, ok := m[n]; ok {
			return true
		}
		m[n] = struct{}{}
	}
	return false
}

func containsDuplicate_anothersolution(nums []int) bool {
	m := make(map[int]struct{})
	for _, n := range nums {
		m[n] = struct{}{}
	}
	return len(nums) > len(m)
}
