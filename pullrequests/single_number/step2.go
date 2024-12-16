//lint:file-ignore U1000 Ignore all unused code
package singlenumber

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
調べたところ下記のようにすれば空間計算量がO(1)になる。
ただし、このXORを使った方法は１つの数以外は全て偶数回必ず出現するという条件を満たさなければ成立しないので、実際にこのように書くのが良い方法だとは思わなかった（ただしビット演算の練習にはなる）。
*/
func singleNumberStep2(nums []int) int {
	singleNum := 0
	for _, n := range nums {
		singleNum ^= n
	}
	return singleNum
}
