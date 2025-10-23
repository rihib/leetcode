//lint:file-ignore U1000 Ignore all unused code
package main

// 愚直にn回掛け算する方法だと、nの最大値は2^31-1で、大体2^10^3なので、
// １ループが数nsだとしても、秒単位のオーダーになってしまい、タイムアウトになる。

// そこで n を2の累乗に分解して計算するバイナリ法と呼ばれる高速化手法が一般に知られている。
// 例えば n=13 は2進数では 1101 である。これは 13 = 8 + 4 + 1 を意味する。
// したがって x^13 = x^8 * x^4 * x となる。
// x, x^2, x^4, x^8 を順に計算していき、n の2進数表現でビットが立っている部分だけを掛け合わせればよい。
// そのため末尾のビットが立っていれば x をかけ、n を右に1ビットシフトするとともに、桁を上げるために x を2乗する。
// 計算量は O(log n) になる。n の2進数表現のビット数に比例するためである。
func myPowBinaryRecursive(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n < 0 {
		return 1 / myPowBinaryRecursive(x, -n)
	}
	powered := 1.0
	if n&0b1 == 1 {
		powered *= x
	}
	powered *= myPowBinaryRecursive(x*x, n>>1)
	return powered
}

func myPowBinaryIterative(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	powered := 1.0
	for n > 0 {
		if n&0b1 == 1 {
			powered *= x
		}
		x *= x
		n >>= 1
	}
	return powered
}

// バイナリ法ではすべてのビットを一つずつチェックところ、sliding window法では複数ビットをひとまとめに見て計算する。
// ひとまとめにするビット幅を窓幅 k とする。窓幅に応じて事前計算を行う。
// 事前計算のコストが増えるが、x が大きい場合にバイナリ法よりも有効である。
// 計算量は事前計算のコストを含めると O((log n)/k + 2^k) になる。
func myPowSlidingWindowK2(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n < 0 {
		return 1.0 / myPowSlidingWindowK2(x, -n)
	}
	powered := 1.0
	switch n & 0b11 {
	case 1:
		powered = x
	case 2:
		powered = x * x
	case 3:
		powered = x * x * x
	}
	powered *= myPowSlidingWindowK2(x*x*x*x, n>>2)
	return powered
}
