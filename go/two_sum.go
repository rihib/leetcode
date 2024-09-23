//lint:file-ignore U1000 Ignore all unused code
package main

func twoSum(nums []int, target int) []int {
	numToIndex := make(map[int]int)
	for i, n := range nums {
		if j, ok := numToIndex[target-n]; ok {
			return []int{i, j}
		}
		numToIndex[n] = i
	}
	return nil // 本来ならerrorを返したい
}
