package queue

type Node struct {
	Val      int
	Priority int
}

type PriorityQueue struct {
	heap []Node
}

func (pq *PriorityQueue) Push(n Node) {
	pq.heap = append(pq.heap, n)
	child := len(pq.heap) - 1
	parent := (child - 1) / 2
	for child > 0 && pq.heap[parent].Priority < n.Priority {
		pq.heap[parent], pq.heap[child] = pq.heap[child], pq.heap[parent]
		child = parent
		parent = (child - 1) / 2
	}
}

func (pq *PriorityQueue) Pop() *Node {
	if len(pq.heap) == 0 {
		return nil
	}

	n := pq.heap[0]
	tail := len(pq.heap) - 1
	pq.heap[0], pq.heap[tail] = pq.heap[tail], pq.heap[0]
	pq.heap = pq.heap[:tail]

	for parent := 0; ; {
		l, r := parent*2+1, parent*2+2
		largest := parent
		if l < len(pq.heap) && pq.heap[l].Priority > pq.heap[parent].Priority {
			largest = l
		}
		if r < len(pq.heap) && pq.heap[r].Priority > pq.heap[parent].Priority {
			largest = r
		}
		if largest == parent {
			break
		}
		pq.heap[parent], pq.heap[largest] = pq.heap[largest], pq.heap[parent]
		parent = largest
	}

	return &n
}

func (pq *PriorityQueue) Peek() *Node {
	if len(pq.heap) == 0 {
		return nil
	}
	return &pq.heap[0]
}

func (pq *PriorityQueue) IsEmpty() bool {
	return len(pq.heap) == 0
}

func (pq *PriorityQueue) Len() int {
	return len(pq.heap)
}
