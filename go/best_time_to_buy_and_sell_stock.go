//lint:file-ignore U1000 Ignore all unused code
package main

func maxProfit(prices []int) int {
	res := 0
	minProfit := prices[0]
	for i := 0; i < len(prices); i++ {
		res = max(res, prices[i]-minProfit)
		minProfit = min(minProfit, prices[i])
	}
	return res
}
