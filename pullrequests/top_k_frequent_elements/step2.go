//lint:file-ignore U1000 Ignore all unused code
package topkfrequentelements

/*
Step1の変数名などをより明確になるように変更しました。
*/
func topKFrequentStep2(nums []int, k int) []int {
	frequency := make(map[int]int)
	for _, num := range nums {
		frequency[num]++
	}
	countToNum := make([][]int, len(nums)+1)
	for num, count := range frequency {
		countToNum[count] = append(countToNum[count], num)
	}
	topK := make([]int, 0, k)
	for i := len(countToNum) - 1; i >= 0 && len(topK) < k; i-- {
		topK = append(topK, countToNum[i]...)
	}
	return topK
}
