package BFS

import "Go-datastructures-algorithm/Graph"

func Bfs(g *Graph.DGraph,start Graph.VertexId,fn func(Graph.VertexId)){
	queue:=[]Graph.VertexId{start}
	visited := make(map[Graph.VertexId]bool)

	var next []Graph.VertexId

	for len(queue)>0{
		next = []Graph.VertexId{}
		for _,vertex:=range queue{
			visited[vertex] = true
			neighbours := g.GetNeighbours(vertex).VerticesIter()
			fn(vertex)

			for neighbour:=range neighbours{
				_,ok:=visited[neighbour]
				if !ok{
					next = append(next,neighbour)
				}
			}
		}
		queue = next
	}
}

