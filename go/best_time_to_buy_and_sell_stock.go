//lint:file-ignore U1000 Ignore all unused code
package main

func maxProfit(prices []int) int {
	minPrice, maxProfit := prices[0], 0
	for _, n := range prices {
		minPrice = min(minPrice, n)
		maxProfit = max(maxProfit, n-minPrice)
	}
	return maxProfit
}
