package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/rihib/leetcode/basics/graph"
	"github.com/rihib/leetcode/basics/queue"
	"github.com/rihib/leetcode/basics/sort"
	"github.com/rihib/leetcode/basics/tree"
)

func main() {
	testSort()
	testPriorityQueue()
	testDijkstra()
	testTree()
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

	fmt.Print("\n")
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

	fmt.Print("\n")
}

func testTree() {
	/*
		下記のツリーが構築される。
								1
						/       \
					2          3
				/   \       / \
			x      4     5   x
						/ \   /
						x x  6
	*/
	nums := []int{1, 2, 3, -1, 4, 5, -1, -1, -1, 6}
	root := buildTree(nums)
	fmt.Println("=====Tree=====")

	fmt.Print("Pre-order: ")
	tree.PreOrderTraversal(root) // 1, 2, 4, 3, 5, 6
	fmt.Print("\n")

	fmt.Print("In-order: ")
	tree.InOrderTraversal(root) // 2, 4, 1, 6, 5, 3
	fmt.Print("\n")

	fmt.Print("Post-order: ")
	tree.PostOrderTraversal(root) // 4, 2, 6, 5, 3, 1
	fmt.Print("\n")

	fmt.Print("Pre-order Iterative: ")
	tree.PreOrderTraversalIterative(root) // 1, 2, 4, 3, 5, 6
	fmt.Print("\n")

	fmt.Print("In-order Iterative: ")
	tree.InOrderTraversalIterative(root) // 2, 4, 1, 6, 5, 3
	fmt.Print("\n")

	fmt.Print("Post-order Iterative: ")
	tree.PostOrderTraversalIterative(root) // 4, 2, 6, 5, 3, 1
	fmt.Print("\n")

	fmt.Print("\n")
}

func buildTree(nums []int) *tree.Node {
	if len(nums) == 0 || nums[0] < 0 {
		return nil
	}
	root := &tree.Node{Val: nums[0]}
	queue := []*tree.Node{root}
	index := 1
	for len(queue) > 0 && index < len(nums) {
		node := queue[0]
		queue = queue[1:]
		if nums[index] >= 0 {
			node.Left = &tree.Node{Val: nums[index]}
			queue = append(queue, node.Left)
		}
		index++
		if index < len(nums) && nums[index] >= 0 {
			node.Right = &tree.Node{Val: nums[index]}
			queue = append(queue, node.Right)
		}
		index++
	}
	return root
}
