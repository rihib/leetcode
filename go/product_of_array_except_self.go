//lint:file-ignore U1000 Ignore all unused code
package main

func productExceptSelf(nums []int) []int {
	allProduct, zeroCount := 0, 0
	for _, n := range nums {
		if n == 0 {
			zeroCount++
			continue
		}
		if allProduct == 0 {
			allProduct = n
		} else {
			allProduct *= n
		}
	}
	products := make([]int, len(nums))
	for i := range products {
		if zeroCount > 0 {
			if nums[i] == 0 && zeroCount == 1 {
				products[i] = allProduct
			} else {
				products[i] = 0
			}
		} else {
			products[i] = allProduct / nums[i]
		}
	}
	return products
}

func productExceptSelf2(nums []int) []int {
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
