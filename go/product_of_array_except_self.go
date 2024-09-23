//lint:file-ignore U1000 Ignore all unused code
package main

func productExceptSelf(nums []int) []int {
	products := make([]int, len(nums))
	leftProduct := 1
	for i := 0; i < len(nums); i++ {
		products[i] = leftProduct
		leftProduct *= nums[i]
	}
	rightProduct := 1
	for i := len(nums) - 1; i >= 0; i-- {
		products[i] *= rightProduct
		rightProduct *= nums[i]
	}
	return products
}
