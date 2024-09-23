//lint:file-ignore U1000 Ignore all unused code
package main

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, n := range nums {
		if j, ok := numsMap[target-n]; ok {
			return []int{i, j}
		}
		numsMap[n] = i
	}
	return nil
}
