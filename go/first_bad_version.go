//lint:file-ignore U1000 Ignore all unused code
package main

const BadVersion = 4

func isBadVersion(version int) bool {
	return version == BadVersion
}

func firstBadVersion(n int) int {
	fbv := n
	h, t := 1, n
	for h <= t {
		mid := (h + t) / 2
		if isBadVersion(mid) {
			fbv = mid
			t = mid - 1
		} else {
			h = mid + 1
		}
	}
	return fbv
}
