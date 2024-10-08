//lint:file-ignore U1000 Ignore all unused code
package main

func climbStairsFibonacci(n int) int {
	prev, curr := 0, 1
	for i := 1; i <= n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}

func climbStairsDP(n int) int {
	m := make(map[int]int, n)
	return climbStairsHelper(n, m)
}

func climbStairsHelper(n int, m map[int]int) int {
	if n == 1 || n == 2 {
		return n
	}
	if ways, ok := m[n]; ok {
		return ways
	}
	m[n] = climbStairsHelper(n-1, m) + climbStairsHelper(n-2, m)
	return m[n]
}
