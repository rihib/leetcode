//lint:file-ignore U1000 Ignore all unused code
package main

type MyQueue struct {
	InputStack  []int
	OutputStack []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.InputStack = append(q.InputStack, x)
}

func (q *MyQueue) Pop() int {
	tail := q.Peek()
	q.OutputStack = q.OutputStack[:len(q.OutputStack)-1]
	return tail
}

func (q *MyQueue) Peek() int {
	if q.Empty() {
		return -1
	}

	if len(q.OutputStack) == 0 {
		for len(q.InputStack) > 0 {
			tail := q.InputStack[len(q.InputStack)-1]
			q.InputStack = q.InputStack[:len(q.InputStack)-1]
			q.OutputStack = append(q.OutputStack, tail)
		}
	}

	return q.OutputStack[len(q.OutputStack)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.InputStack) == 0 && len(q.OutputStack) == 0
}
