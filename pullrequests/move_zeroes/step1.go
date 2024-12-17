//lint:file-ignore U1000 Ignore all unused code
package movezeroes

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
左側に非ゼロのエリア、右側にゼロのエリアを作れば良い。
イメージとしてはzeroIndexはゼロのエリアのはじまり（境界）のインデックスを持ち、
for文で探索していって非ゼロの要素を見つけたらzeroIndexの要素とスワップして、
zeroIndexをインクリメントすることで、zeroIndexの左側に非ゼロの要素を集め、
右側にゼロを集めるという感じ。

境界というのを明確にしたいのであればstartZeroIndexとかでもいいかも。またはzeroIndexだとインデックスの値自体が0なのかという誤解を引き起こす恐れがあるので、indexOfStartZeroIndexとかでもいいかもしれない。ただ全体の行数に対して変数名が冗長すぎる気がしたので今回はzeroIndexにした。
*/
func moveZeroes(nums []int) {
	zeroIndex := 0
	for i, n := range nums {
		if n == 0 {
			continue
		}
		nums[zeroIndex], nums[i] = nums[i], nums[zeroIndex]
		zeroIndex++
	}
}
