//lint:file-ignore U1000 Ignore all unused code
package majorityelement

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
時間：3分

一番自然な方法で解くことにした。
最後にLeetCodeの制約上しょうがなく-1を返しているが、本当はerrorを返したい。
*/
func majorityElement(nums []int) int {
	frequencies := make(map[int]int, len(nums))
	for _, n := range nums {
		frequencies[n]++
		if frequencies[n] > len(nums)/2 {
			return n
		}
	}
	return -1
}
