//lint:file-ignore U1000 Ignore all unused code
package main

func sortedSquares(nums []int) []int {
	// [-4,-1,0,3,10]
	startIdx := 0
	for ; startIdx < len(nums); startIdx++ {
		if nums[startIdx] >= 0 {
			break
		}
	}

	var sorted []int
	for l, r := startIdx-1, startIdx; 0 <= l || r < len(nums); {
		if l < 0 {
			sorted = append(sorted, nums[r]*nums[r])
			r++
			continue
		}
		if len(nums) <= r {
			sorted = append(sorted, nums[l]*nums[l])
			l--
			continue
		}

		if nums[l]*nums[l] <= nums[r]*nums[r] {
			sorted = append(sorted, nums[l]*nums[l])
			l--
		} else {
			sorted = append(sorted, nums[r]*nums[r])
			r++
		}
	}
	return sorted
}
