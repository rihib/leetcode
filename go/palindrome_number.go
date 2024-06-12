//lint:file-ignore U1000 Ignore all unused code
package main

func isPalindromeNumber(x int) bool {
	// TestCase: -1, 0, 1, 12, 121, 1231

	if x < 0 {
		return false
	}
	if x <= 9 {
		return true
	}

	var nums []int
	for x > 0 {
		nums = append(nums, x%10)
		x /= 10
	}

	for i, n := range nums {
		j := len(nums) - 1 - i

		if i >= j {
			break
		}

		if n != nums[j] {
			return false
		}
	}

	return true
}
