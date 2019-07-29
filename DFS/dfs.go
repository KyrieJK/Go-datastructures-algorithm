package DFS

import "Go-datastructures-algorithm/Graph"
import "Go-datastructures-algorithm/Stack"

func UndirectedDfs(g *Graph.UndGraph, v Graph.VertexId, fn func(Graph.VertexId)) {
	s := Stack.New()
	s.Push(v)
	visited := make(map[Graph.VertexId]bool)

	for s.Len() > 0 {
		v:=s.Pop().(Graph.VertexId)

		if _,ok:=visited[v];!ok{
			visited[v] = true
			fn(v)
			neighbours := g.GetNeighbours(v).VerticesIter()
			for neighbour:=range neighbours{
				s.Push(neighbour)
			}
		}
	}
}

func DirectedDfs(g *Graph.DGraph,v Graph.VertexId,fn func(Graph.VertexId)){
	s := Stack.New()
	s.Push(v)
	visited := make(map[Graph.VertexId]bool)

	for s.Len() >0{
		v = s.Pop().(Graph.VertexId)

		if _,ok:=visited[v];!ok{
			visited[v] = true
			fn(v)
			neighbours:=g.GetSuccessors(v).VerticesIter()
			for neighbour := range neighbours{
				s.Push(neighbour)
			}
		}
	}
}
