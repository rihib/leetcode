//lint:file-ignore U1000 Ignore all unused code
package main

func maxProfit(prices []int) int {
	minPrice, maxProfit := prices[0], 0
	for _, price := range prices {
		minPrice = min(minPrice, price)
		maxProfit = max(maxProfit, price-minPrice)
	}
	return maxProfit
}
