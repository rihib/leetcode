//lint:file-ignore U1000 Ignore all unused code
package main

func singleNumber(nums []int) int {
	res := 0
	for _, n := range nums {
		res ^= n
	}
	return res
}
