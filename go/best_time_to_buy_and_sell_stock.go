package blind75

import "math"

func maxProfit(prices []int) int {
	res := 0
	minPrice := math.MaxUint32

	for _, price := range prices {
		minPrice = min(price, minPrice)
		res = max(price-minPrice, res)
	}

	return res
}
