package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/rihib/leetcode/basics/sort"
)

func main() {
	testcases := [][]int{
		{},
		{1},
		{1, 1},
		{3, 2, 1},
		{3, 2, 2, 1},
		{4, 3, 2, 1},
		{4, 4, 3, 3, 2, 2, 1, 1},
	}

	runSort(sort.Insertion, testcases)
	runSort(sort.Selection, testcases)
	runSort(sort.Quicksort, testcases)
	runSort(sort.Merge, testcases)
}

func runSort(f func([]int) []int, testcases [][]int) {
	fmt.Println(getFunctionName(f))
	for _, nums := range testcases {
		sorted := f(nums)
		print(sorted)
	}
	fmt.Print("\n")
}

func getFunctionName(f interface{}) string {
	full := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	n := "=====" +
		strings.TrimPrefix(full, "github.com/rihib/leetcode/basics/sort.") +
		"====="
	return n
}

func print(nums []int) {
	if len(nums) == 0 {
		fmt.Print("Empty")
	}
	for _, n := range nums {
		fmt.Printf("%d ", n)
	}
	fmt.Print("\n")
}
