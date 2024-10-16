//lint:file-ignore U1000 Ignore all unused code
package main

// 2の冪乗の時はMSBのみが1になる
// それ以外の場合はMSB以外の部分はすでに前に計算していてメモ化されているのでそれを使えばいい
func countBits(n int) []int {
	bitCounts := make([]int, n+1)
	powerOfTwo := 1
	for i := 1; i <= n; i++ {
		if i == powerOfTwo*2 {
			powerOfTwo = i
		}
		bitCounts[i] = 1 + bitCounts[i-powerOfTwo]
	}
	return bitCounts
}

// 内側のループはnを2で割り続けるのでlog n回
// 外側のループはn回なので、全体でO(n log n)になる
func countBits2(n int) []int {
	bitCounts := make([]int, n+1)
	for i := 0; i <= n; i++ {
		num := i
		count := 0
		for num > 0 {
			count += num % 2
			num /= 2
		}
		bitCounts[i] = count
	}
	return bitCounts
}
