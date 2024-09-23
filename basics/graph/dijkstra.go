package graph

import (
	"container/heap"
	"math"
)

type Edge struct {
	To     int
	Weight int
}

// adjacency list
type Graph struct {
	Nodes map[int][]Edge
}

func (g *Graph) Dijkstra(source int) map[int]int {
	dist := make(map[int]int)
	for node := range g.Nodes {
		dist[node] = math.MaxInt64
	}
	dist[source] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{Node: source, Priority: 0})

	for pq.Len() > 0 {
		u := heap.Pop(&pq).(*Item).Node

		for _, edge := range g.Nodes[u] {
			v := edge.To
			weight := edge.Weight
			if dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
				heap.Push(&pq, &Item{Node: v, Priority: dist[v]})
			}
		}
	}

	return dist
}
