//lint:file-ignore U1000 Ignore all unused code
package main

import "math"

func maxProfit(prices []int) int {
	minPrice, maxValue := math.MaxInt, 0
	for _, price := range prices {
		minPrice = min(minPrice, price)
		maxValue = max(maxValue, price-minPrice)
	}
	return maxValue
}
