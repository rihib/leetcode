//lint:file-ignore U1000 Ignore all unused code
package main

func longestConsecutive(nums []int) int {
	numsMap := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		numsMap[n] = struct{}{}
	}
	maxLength := 0
	for _, n := range nums {
		if _, ok := numsMap[n-1]; ok {
			continue
		}
		length := 1
		for {
			n++
			if _, ok := numsMap[n]; !ok {
				break
			}
			length++
		}
		maxLength = max(maxLength, length)
	}
	return maxLength
}
