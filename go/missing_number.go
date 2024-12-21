//lint:file-ignore U1000 Ignore all unused code
package main

func missingNumber(nums []int) int {
	var difference int
	for i, n := range nums {
		difference += i + 1 - n
	}
	return difference
}

func missingNumber2(nums []int) int {
	var total int
	for _, n := range nums {
		total += n
	}
	return (1+len(nums))*len(nums)/2 - total
}
