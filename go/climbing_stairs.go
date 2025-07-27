//lint:file-ignore U1000 Ignore all unused code
package main

func climbStairsIterative(n int) int {
	current, next := 1, 1
	for range n {
		current, next = next, current+next
	}
	return current
}

func climbStairsRecursive(n int) int {
	memo := make(map[int]int, n+1)
	return climbStairsWithMemo(n, memo)
}

func climbStairsWithMemo(n int, memo map[int]int) int {
	if n == 1 || n == 2 {
		return n
	}
	if steps, ok := memo[n]; ok {
		return steps
	}
	memo[n] = climbStairsWithMemo(n-1, memo) + climbStairsWithMemo(n-2, memo)
	return memo[n]
}
