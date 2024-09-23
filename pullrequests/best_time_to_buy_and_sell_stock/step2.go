//lint:file-ignore U1000 Ignore all unused code
package besttimetobuyandsellstock

import "math"

/*
関数名と被らないように変数名を更新したのと、minPriceをmath.MaxIntで初期化することで、たとえ空の入力が渡されたとしてもエラーにならないようにした。
*/
func maxProfit(prices []int) int {
	minPrice, maxValue := math.MaxInt, 0
	for _, price := range prices {
		minPrice = min(minPrice, price)
		maxValue = max(maxValue, price-minPrice)
	}
	return maxValue
}
