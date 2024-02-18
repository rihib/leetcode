//lint:file-ignore U1000 Ignore all unused code
package main

import "math"

func maxProfit(prices []int) int {
	res, minProfit := 0, math.MaxUint32
	for _, price := range prices {
		res, minProfit = max(res, price-minProfit), min(minProfit, price)
	}
	return res
}
