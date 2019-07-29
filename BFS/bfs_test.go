package BFS

import(
	"fmt"
	"Go-datastructures-algorithm/Graph"
	"testing"
)

func TestBfs(t *testing.T){
	h:=Graph.NewDirected()

	for i := 0; i < 10; i++ {
		v := Graph.VertexId(i)
		h.AddVertex(v)
	}

	for i := 0; i < 9; i++ {
		h.AddEdge(Graph.VertexId(i),Graph.VertexId(i+1),1)
	}

	bfsMap:=make(map[Graph.VertexId]bool)
	checkVertices := func(v Graph.VertexId){
		bfsMap[v] = true
	}

	Bfs(h,Graph.VertexId(3),checkVertices)

	for i := 3; i < len(bfsMap); i++ {
		if _,ok := bfsMap[Graph.VertexId(i)];!ok{
			fmt.Println(bfsMap)
			t.Error()
		}
	}
}

