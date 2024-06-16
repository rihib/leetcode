//lint:file-ignore U1000 Ignore all unused code
package main

func maxProfit(prices []int) int {
	lowestPrice, highestProfit := prices[0], 0
	for _, n := range prices {
		lowestPrice = min(lowestPrice, n)
		highestProfit = max(highestProfit, n-lowestPrice)
	}
	return highestProfit
}
