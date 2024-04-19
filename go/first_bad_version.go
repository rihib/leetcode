//lint:file-ignore U1000 Ignore all unused code
package main

const BadVersion = 4

func isBadVersion(version int) bool {
	return version == BadVersion
}

func firstBadVersion(n int) int {
	l, r := 1, n
	for l < r {
		mid := (l + r) / 2
		if isBadVersion(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
