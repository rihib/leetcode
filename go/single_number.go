//lint:file-ignore U1000 Ignore all unused code
package main

// nums = [2, 3, 2]のとき
// 1回目のループ：0000 XOR 0010 = 0010
// 2回目のループ：0010 XOR 0011 = 0001
// 3回目のループ：0001 XOR 0010 = 0011
// つまり１つだけ１回だけ出現し、他は全て２回出現するという条件が満たされない時は
// XORを使った方法はうまくいかない（例えば[2, 3]など）
func singleNumber(nums []int) int {
	singleNum := 0
	for _, n := range nums {
		singleNum ^= n
	}
	return singleNum
}
