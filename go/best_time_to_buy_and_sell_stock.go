package blind75

import "math"

func maxProfit(prices []int) int {
	res := 0
	minProfit := math.MaxUint32
	for _, price := range prices {
		res = max(res, price-minProfit)
		minProfit = min(minProfit, price)
	}
	return res
}
