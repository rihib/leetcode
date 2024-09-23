//lint:file-ignore U1000 Ignore all unused code
package template

/*
入力にアルファベットの小文字しかないので配列でも実装してみた。
mapを使った方法では実行時間が30~35msほどだったのに対し、配列を使う方法では10msほどである。
この理由として主にキャッシュ効率とオーバーヘッドの違いが考えられる。

配列は要素が連続したメモリ領域に配置されるのに対し、mapでは内部的には８つのエントリを持つバケットから構成されていて、各バケットはメモリ上で異なる位置に存在するため、参照局所性が悪い。
また配列では要素が26個しかないことからサイズは26*8B=208Bより、十分1次キャッシュに載るのに対し、mapは１エントリあたり4B+8B=12Bで、最大サイズは12B*10^5=1.2GBになることから３次キャッシュにも載り切らないというのもある。

また、mapはハッシュ化のオーバーヘッドがあるのと、Step1の場合はcapacityを指定せずにmapを作成したため、Load factorがある閾値（6.5)を超えるとリサイズが行われ、リハッシュする必要が生じる。
よって主に上記の２つを要因としてmapの方が実行時間が遅くなったと考えられる。
ちなみにStep2ではcapacityを指定してmapを作成したところ、実行時間は20~25msとなった。
*/
func firstUniqCharStep2(s string) int {
	frequency := make(map[rune]int, len(s))
	for _, r := range s {
		frequency[r]++
	}
	for i, r := range s {
		if frequency[r] == 1 {
			return i
		}
	}
	return -1
}

func firstUniqChar2Step2(s string) int {
	var frequency [26]int
	for _, r := range s {
		frequency[r-'a']++
	}
	for i, r := range s {
		if frequency[r-'a'] == 1 {
			return i
		}
	}
	return -1
}
