//lint:file-ignore U1000 Ignore all unused code
package singlenumber

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
時間：３分
LeetCodeには空間計算量がO(1)になるように実装しろと書かれていたが、方法が思いつかなかったので、まずは普通に実装してみた。
*/
func singleNumberStep1(nums []int) int {
	singleNums := make(map[int]int, len(nums))
	for _, n := range nums {
		singleNums[n]++
		if singleNums[n] > 1 {
			delete(singleNums, n)
		}
	}
	// 必要に応じてlen(singleNums) != 1のときはエラーを返すという処理も追加しても良いかも
	for n := range singleNums {
		return n
	}
	return 0 // 本来は見つからなかった場合はエラーを返したい
}
