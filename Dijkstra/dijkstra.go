package Dijkstra

import (
	"Go-datastructures-algorithm/Graph"
	"Go-datastructures-algorithm/Priority_Queue"
)

func Dijkstra(g *Graph.UndGraph, source Graph.VertexId) map[Graph.VertexId]Graph.VertexId {
	visited := make(map[Graph.VertexId]bool, g.VerticesCount())
	dist := make(map[Graph.VertexId]int)
	prev := make(map[Graph.VertexId]Graph.VertexId)
	Q := Priority_Queue.NewMin() //创建最小堆，堆顶为最小元素

	vertices := g.VerticesIter()

	dist[source] = 0
	for vertex := range vertices {
		if vertex != source {
			dist[vertex] = 1000
			prev[vertex] = 0
		}
		Q.Insert(*Priority_Queue.NewItem(vertex, dist[vertex]))
	}

	for Q.Len() > 0 {
		u := Q.Extract().Value.(Graph.VertexId)
		visited[u] = true

		for neighbour := range g.GetNeighbours(u).VerticesIter() {
			if !visited[neighbour] {
				alt := dist[u] + g.GetEdge(u, neighbour)

				if alt < dist[neighbour] {
					dist[neighbour] = alt
					prev[neighbour] = u
					Q.ChangePriority(neighbour, alt)
				}
			}
		}
	}
	return prev
}
