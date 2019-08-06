package Dijkstra

import (
	"Go-datastructures-algorithm/Graph"
	"fmt"
	"testing"
)

func TestDijkstra(t *testing.T) {
	h := Graph.NewUndirected()

	for i := 0; i < 5; i++ {
		h.AddVertex(Graph.VertexId(i))
	}

	h.AddEdge(Graph.VertexId(0), Graph.VertexId(1), 10)
	h.AddEdge(Graph.VertexId(1), Graph.VertexId(2), 20)
	h.AddEdge(Graph.VertexId(2), Graph.VertexId(3), 40)
	h.AddEdge(Graph.VertexId(0), Graph.VertexId(2), 50)
	h.AddEdge(Graph.VertexId(0), Graph.VertexId(3), 80)
	h.AddEdge(Graph.VertexId(0), Graph.VertexId(4), 10)
	h.AddEdge(Graph.VertexId(4), Graph.VertexId(3), 10)

	prev := Dijkstra(h, Graph.VertexId(0))

	if prev[1] != Graph.VertexId(0) || prev[2] != Graph.VertexId(1) || prev[3] != Graph.VertexId(4) || prev[4] != Graph.VertexId(0) {
		fmt.Println(prev)
		t.Error()
	}
}
