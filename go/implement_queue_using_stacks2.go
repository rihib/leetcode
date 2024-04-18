//lint:file-ignore U1000 Ignore all unused code
package main

type MyQueue2 struct {
	Queue []int
}

func Constructor2() MyQueue2 {
	return MyQueue2{}
}

func (q *MyQueue2) Push(x int) {
	q.Queue = append(q.Queue, x)
}

func (q *MyQueue2) Pop() int {
	x := q.Queue[0]
	q.Queue = q.Queue[1:]
	return x
}

func (q *MyQueue2) Peek() int {
	return q.Queue[0]
}

func (q *MyQueue2) Empty() bool {
	return len(q.Queue) == 0
}
