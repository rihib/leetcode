//lint:file-ignore U1000 Ignore all unused code
package main

func reverseBits(nums uint32) uint32 {
	var reversed uint32
	for i := 0; i < 32; i++ {
		reversed = reversed<<1 + nums%2
		nums >>= 1
	}
	return reversed
}

func reverseBits2(num uint32) uint32 {
	var res uint32 = 0
	for i := 0; i < 32; i++ {
		bit := (num >> i) & 1
		res = res | (bit << (31 - i))
	}
	return res
}
