package BFS

import (
	"Go-datastructures-algorithm/Graph"
)

func ShortestPath(g *Graph.DGraph, start Graph.VertexId) (dist map[Graph.VertexId]int) {
	dist = make(map[Graph.VertexId]int)
	visited := make(map[Graph.VertexId]bool)

	getDist := func(v Graph.VertexId) {
		neighbours := g.GetNeighbours(v).VerticesIter()
		visited[v] = true

		for neighbour := range neighbours {
			ok, _ := visited[neighbour]
			if !ok {
				dist[neighbour] = dist[v] + 1
			}
		}
	}

	
	queue := []Graph.VertexId{start}

	var next []Graph.VertexId

	for len(queue) > 0 {
		next = []Graph.VertexId{}
		for _, vertex := range queue {
			visited[vertex] = true
			neighbours := g.GetNeighbours(vertex).VerticesIter()
			getDist(vertex)

			for neighbour := range neighbours {
				_, ok := visited[neighbour]
				if !ok {
					next = append(next, neighbour)
				}
			}
		}
		queue = next
	}

	return dist
}

func GetDist(g *Graph.DGraph, from Graph.VertexId, to Graph.VertexId) int {
	return ShortestPath(g, from)[to]
}
