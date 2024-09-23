//lint:file-ignore U1000 Ignore all unused code
package longestconsecutivesequence

/*
他の人のPRを見て、「手でやるとして、{0, 1, 2, 3, 4, 5} だったとして、0 から確認が終わった後に 1 を確認しないじゃないですか」（https://github.com/t-ooka/leetcode/pull/7/files#r1665902244 ）というodaさんのコメントを発見し、確かにその通りだと思ったので下記のように一度連続する数に含めたものは削除するように変更した。
実際にLeetCode上でもStep1の場合、５回ほど提出しても実行時間がどれも1800msほどだったのに対し、Step2の場合は90msぐらいだったので、実際にとても早くなることが実感できた。

質問：どうやらこの問題はUnion FIndを使っても解くことができるようですが、Union FIndを使って解く練習もするべきでしょうか？
*/
func longestConsecutiveStep2(nums []int) int {
	numsMap := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		numsMap[n] = struct{}{}
	}
	maxLength := 0
	for n := range numsMap {
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
			delete(numsMap, n)
		}
		maxLength = max(maxLength, length)
	}
	return maxLength
}
