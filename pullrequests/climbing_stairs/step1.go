//lint:file-ignore U1000 Ignore all unused code
package template

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
時間：11分

フィボナッチだとわかったので、その方法で解いた。
少し頭がこんがらがってしまったので時間がかかってしまった。
本来はn <= 0のときはエラーにするなどすべきだと思う。
*/
func climbStairsFibonacci(n int) int {
	prev, curr := 0, 1
	for i := 1; i <= n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}
