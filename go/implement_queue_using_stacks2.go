//lint:file-ignore U1000 Ignore all unused code
package main

import "fmt"

type Queue struct {
	pushStack []int
	popStack  []int
}

func (q *Queue) Push(n int) {
	q.pushStack = append(q.pushStack, n)
}

func (q *Queue) Pop() (int, error) {
	n, err := q.Peek()
	if err != nil {
		return 0, fmt.Errorf("queue is empty, cannot pop")
	}
	q.popStack = q.popStack[:len(q.popStack)-1]
	return n, nil
}

func (q *Queue) Peek() (int, error) {
	if q.Empty() {
		return 0, fmt.Errorf("queue is empty, cannot peek")
	}
	if len(q.popStack) == 0 {
		for len(q.pushStack) > 0 {
			q.popStack = append(q.popStack, q.pushStack[len(q.pushStack)-1])
			q.pushStack = q.pushStack[:len(q.pushStack)-1]
		}
	}
	return q.popStack[len(q.popStack)-1], nil
}

func (q *Queue) Empty() bool {
	return len(q.pushStack) == 0 && len(q.popStack) == 0
}
