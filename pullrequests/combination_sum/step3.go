//lint:file-ignore U1000 Ignore all unused code
package template

/*
動的計画法を使った方法も実装しました。
targetごとに組み合わせを求めると重複する組み合わせが生じてしまうので、candidateごとに組み合わせを求めるようにしました。
*/
func combinationSum_dp(candidates []int, target int) [][]int {
	combinationsGroups := make([][][]int, target+1)
	combinationsGroups[0] = [][]int{{}}
	for _, candidate := range candidates {
		for i := candidate; i <= target; i++ {
			for _, combination := range combinationsGroups[i-candidate] {
				newCombination := append([]int{}, combination...)
				newCombination = append(newCombination, candidate)
				combinationsGroups[i] = append(combinationsGroups[i], newCombination)
			}
		}
	}
	return combinationsGroups[target]
}
