package main

import (
	"fmt"

	"github.com/rihib/leetcode/basics/sort"
)

func main() {
	nums := []int{4, 3, 2, 1}
	sortedNums := sort.Insertion(nums)
	for _, n := range sortedNums {
		fmt.Println(n)
	}
}
