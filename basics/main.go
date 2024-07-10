package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/rihib/leetcode/basics/graph"
	"github.com/rihib/leetcode/basics/queue"
	"github.com/rihib/leetcode/basics/sort"
)

func main() {
	testSort()
	testPriorityQueue()
	testDijkstra()
}

func testSort() {
	testcases := [][]int{
		{},
		{1},
		{1, 1},
		{3, 2, 1},
		{3, 2, 2, 1},
		{4, 3, 2, 1},
		{1, 2, 3, 4},
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

func testPriorityQueue() {
	fmt.Println("=====Priority Queue=====")
	pq := new(queue.PriorityQueue)

	node, l, isEmpty := pq.Peek(), pq.Len(), pq.IsEmpty()
	fmt.Printf("Peek returns nil: %t, Len: %d, IsEmpty: %t\n", node == nil, l, isEmpty)

	node = pq.Pop()
	fmt.Printf("Pop returns nil: %t\n", node == nil)

	node = &queue.Node{Priority: 100, Val: 1}
	fmt.Printf("Push Node: Priority is %d, Val is %d\n", node.Priority, node.Val)
	pq.Push(*node)

	node, l, isEmpty = pq.Peek(), pq.Len(), pq.IsEmpty()
	fmt.Printf("Peek: Priority is %d and Val is %d, Len: %d, IsEmpty: %t\n", node.Priority, node.Val, l, isEmpty)

	node = pq.Pop()
	fmt.Printf("Pop Node: Priority is %d, Val is %d\n", node.Priority, node.Val)

	node, l, isEmpty = pq.Peek(), pq.Len(), pq.IsEmpty()
	fmt.Printf("Peek returns nil: %t, Len: %d, IsEmpty: %t\n", node == nil, l, isEmpty)

	node = &queue.Node{Priority: 100, Val: 1}
	fmt.Printf("Push Node: Priority is %d, Val is %d\n", node.Priority, node.Val)
	pq.Push(*node)

	node = &queue.Node{Priority: 200, Val: 2}
	fmt.Printf("Push Node: Priority is %d, Val is %d\n", node.Priority, node.Val)
	pq.Push(*node)

	node = &queue.Node{Priority: 300, Val: 3}
	fmt.Printf("Push Node: Priority is %d, Val is %d\n", node.Priority, node.Val)
	pq.Push(*node)

	node, l, isEmpty = pq.Peek(), pq.Len(), pq.IsEmpty()
	fmt.Printf("Peek: Priority is %d and Val is %d, Len: %d, IsEmpty: %t\n", node.Priority, node.Val, l, isEmpty)

	node = pq.Pop()
	fmt.Printf("Pop Node: Priority is %d, Val is %d\n", node.Priority, node.Val)

	node = pq.Pop()
	fmt.Printf("Pop Node: Priority is %d, Val is %d\n", node.Priority, node.Val)

	node = pq.Pop()
	fmt.Printf("Pop Node: Priority is %d, Val is %d\n", node.Priority, node.Val)

	node, l, isEmpty = pq.Peek(), pq.Len(), pq.IsEmpty()
	fmt.Printf("Peek returns nil: %t, Len: %d, IsEmpty: %t\n", node == nil, l, isEmpty)

	fmt.Println("")
}

func testDijkstra() {
	fmt.Println("=====Dijkstra=====")
	g := &graph.Graph{
		Nodes: make(map[int][]graph.Edge),
	}
	g.Nodes[0] = []graph.Edge{{To: 1, Weight: 4}, {To: 2, Weight: 1}}
	g.Nodes[1] = []graph.Edge{{To: 3, Weight: 1}}
	g.Nodes[2] = []graph.Edge{{To: 1, Weight: 2}, {To: 3, Weight: 5}}
	g.Nodes[3] = []graph.Edge{}

	dist := g.Dijkstra(0)
	for node, distance := range dist {
		fmt.Printf("Distance from node 0 to node %d is %d\n", node, distance)
	}
}
