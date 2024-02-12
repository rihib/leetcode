package blind75

import "math"

func maxProfit(prices []int) int {
	res := 0
	min := math.MaxUint32

	for _, price := range prices {
		if price <= min {
			min = price
			continue
		}

		if price-min > res {
			res = price - min
		}
	}

	return res
}
