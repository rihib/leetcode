//lint:file-ignore U1000 Ignore all unused code
package main

const BadVersion = 4

func isBadVersion(version int) bool {
	return version == BadVersion
}

func firstBadVersionHalfClosed(n int) int {
	left, right := 1, n+1
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func firstBadVersionClosed(n int) int {
	left, right := 1, n
	for left <= right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
