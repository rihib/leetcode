//lint:file-ignore U1000 Ignore all unused code
package main

// `implement_queue_using_stacks2.go`に実際に書く場合のコードを記載
type MyQueue struct {
	pushStack []int
	popStack  []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(n int) {
	q.pushStack = append(q.pushStack, n)
}

func (q *MyQueue) Pop() int {
	n := q.Peek()
	q.popStack = q.popStack[:len(q.popStack)-1]
	return n
}

func (q *MyQueue) Peek() int {
	if len(q.popStack) > 0 {
		return q.popStack[len(q.popStack)-1]
	}
	for len(q.pushStack) > 0 {
		q.popStack = append(q.popStack, q.pushStack[len(q.pushStack)-1])
		q.pushStack = q.pushStack[:len(q.pushStack)-1]
	}
	return q.popStack[len(q.popStack)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.pushStack) == 0 && len(q.popStack) == 0
}
