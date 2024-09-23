//lint:file-ignore U1000 Ignore all unused code
package template

/*
時間：６分
全体の積を求め、個々の数字で割ることで求める。
解いた後に割り算を使ってはいけないということに気づいた。
*/
func productExceptSelfStep1(nums []int) []int {
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
