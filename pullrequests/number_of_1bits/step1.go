//lint:file-ignore U1000 Ignore all unused code
package numberof1bits

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
時間：3分

rightmost set bitは n & (n - 1) で求めることができるので、効率的にcountしていくことができる。
*/
func hammingWeightStep1(n int) int {
	count := 0
	for n > 0 {
		n &= n - 1
		count += 1
	}
	return count
}
