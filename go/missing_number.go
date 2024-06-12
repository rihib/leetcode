//lint:file-ignore U1000 Ignore all unused code
package main

func missingNumber(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	l := len(nums)
	return l*(l+1)/2 - sum
}

func missingNumber2(nums []int) int {
	res := len(nums)
	for i, n := range nums {
		res += i - n
	}
	return res
}
