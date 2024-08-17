//lint:file-ignore U1000 Ignore all unused code
package main

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	lp := 1
	for i := 0; i < len(nums); i++ {
		res[i] = lp
		lp *= nums[i]
	}

	rp := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= rp
		rp *= nums[i]
	}

	return res
}

func productExceptSelf2(nums []int) []int {
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
	for i, _ := range products {
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
