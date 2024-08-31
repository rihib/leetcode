//lint:file-ignore U1000 Ignore all unused code
package besttimetobuyandsellstock

/*
時間：４分30秒

走査している中で現在の値に対して最大の利益を取り得るのは今まで走査してきた中の最小値であるため、最小値を更新した後に、現在の値に対する利益を計算し、それが最大値であれば更新するようにしている。

解いた後に中の変数名と関数名が一緒になってしまっていることに気づいた。
*/
func maxProfitStep1(prices []int) int {
	minPrice, maxValue := prices[0], 0
	for _, price := range prices {
		minPrice = min(minPrice, price)
		maxValue = max(maxValue, price-minPrice)
	}
	return maxValue
}
