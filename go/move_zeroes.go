//lint:file-ignore U1000 Ignore all unused code
package main

// 左側に非ゼロのエリア、右側にゼロのエリアを作れば良い。
// イメージとしてはzeroIndexはゼロのエリアのはじまり（境界）のインデックスを持ち、
// for文で探索していって非ゼロの要素を見つけたらzeroIndexの要素とスワップして、
// zeroIndexをインクリメントすることで、zeroIndexの左側に非ゼロの要素を集め、
// 右側にゼロを集めるという感じ
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
