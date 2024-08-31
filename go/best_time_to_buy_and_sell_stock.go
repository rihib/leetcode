//lint:file-ignore U1000 Ignore all unused code
package main

func maxProfit(prices []int) int {
	minPrice, maxValue := prices[0], 0
	for _, price := range prices {
		minPrice = min(minPrice, price)
		maxValue = max(maxValue, price-minPrice)
	}
	return maxValue
}
