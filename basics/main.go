package main

import (
	"fmt"

	"github.com/rihib/leetcode/basics/sort"
)

func main() {
	nums := []int{4, 3, 2, 1}
	sorted := sort.Insertion(nums)
	print(sorted)
}

func print(nums []int) {
	for _, n := range nums {
		fmt.Printf("%d ", n)
	}
	fmt.Print("\n")
}
