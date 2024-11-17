//lint:file-ignore U1000 Ignore all unused code
package main

func climbStairsDP(n int) int {
	prev, curr := 1, 1
	for i := 1; i < n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}

func climbStairsmemorizedRecursion(n int) int {
	m := make(map[int]int, n)
	return climbStairsHelper(n, m)
}

func climbStairsHelper(n int, m map[int]int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if ways, ok := m[n]; ok {
		return ways
	}
	m[n] = climbStairsHelper(n-1, m) + climbStairsHelper(n-2, m)
	return m[n]
}
