//lint:file-ignore U1000 Ignore all unused code
package longestconsecutivesequence

/*
時間：9分30秒

最初にソートしようかと思ったが、問題の制約にO(n)で解くことという制約があったので、こういう時は大体配列をマップにする前処理をすれば良いことが多い気がしたので、そこから下記の方法を思いついた。
一見すると連続する数が多いほどO(n^2)に近づく気がするが、連続するほど他の数は飛ばされることが多くなるわけなので、O(n)で解くことができる。
*/
func longestConsecutiveStep1(nums []int) int {
	numsMap := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		numsMap[n] = struct{}{}
	}
	maxLength := 0
	for _, n := range nums {
		if _, ok := numsMap[n-1]; ok {
			continue
		}
		length := 1
		for {
			n++
			if _, ok := numsMap[n]; !ok {
				break
			}
			length++
		}
		maxLength = max(maxLength, length)
	}
	return maxLength
}
