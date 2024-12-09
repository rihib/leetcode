//lint:file-ignore U1000 Ignore all unused code
package numberof1bits

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
ナイーブに解くこともできる。この場合はO(n)の時間計算量になり、効率が悪い。
*/
func hammingWeightStep3_1(n int) int {
	count := 0
	for n > 0 {
		count += n % 2
		n /= 2
	}
	return count
}

func hammingWeightStep3_2(n int) int {
	count := 0
	for n != 0 {
		count += n & 1
		n >>= 1
	}
	return count
}
