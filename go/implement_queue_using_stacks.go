//lint:file-ignore U1000 Ignore all unused code
package main

type MyQueue struct {
	Queue []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.Queue = append(q.Queue, x)
}

func (q *MyQueue) Pop() int {
	x := q.Queue[0]
	q.Queue = q.Queue[1:]
	return x
}

func (q *MyQueue) Peek() int {
	return q.Queue[0]
}

func (q *MyQueue) Empty() bool {
	return len(q.Queue) == 0
}
