package blind75

import "math"

func maxProfit(prices []int) int {
	maxProfit := 0
	min := math.MaxUint32

	for _, price := range prices {
		if price <= min {
			min = price
			continue
		}

		if price-min > maxProfit {
			maxProfit = price - min
		}
	}

	return maxProfit
}
