//lint:file-ignore U1000 Ignore all unused code
package main

func singleNumber(nums []int) int {
	singleNum := 0
	for _, n := range nums {
		singleNum ^= n
	}
	return singleNum
}
