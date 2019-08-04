package BFS

import (
	"Go-datastructures-algorithm/Graph"
	"testing"
)

func TestBfsShortestPath(t *testing.T) {
	h := Graph.NewDirected()

	for i := 0; i < 10; i++ {
		v := Graph.VertexId(i)
		h.AddVertex(v)
	}

	for i := 0; i < 9; i++ {
		h.AddEdge(Graph.VertexId(i), Graph.VertexId(i+1), 1)
	}

	distance := ShortestPath(h, Graph.VertexId(0))

	for i := 0; i < len(distance); i++ {
		if distance[Graph.VertexId(i)] != i {
			t.Error()
		}

		if GetDist(h, Graph.VertexId(0), Graph.VertexId(i)) != i {
			t.Error()
		}
	}
}
