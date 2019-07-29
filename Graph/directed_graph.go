package Graph

type DGraph struct{
	graph
}

func NewDirected() *DGraph{
	return &DGraph{
		graph{
			edgesCount:0,
			edges:make(map[VertexId]map[VertexId]int),
			isDirected:true,
		},
	}
}

func (g *graph) GetPredecessors(vertex VertexId) VerticesIterable{
	iterator := func() <-chan VertexId{
		ch := make(chan VertexId)
		go func(){
			if connected,ok:=g.edges[vertex];ok{
				for vertexId,_:=range connected{
					if g.CheckEdge(vertexId,vertex){
						ch <- vertexId
					}
				}
			}
			close(ch)
		}()
		return ch
	}

	return VerticesIterable(&vertexIterableHelper{iterFunc:iterator})
}

func (g *graph) GetSuccessors(vertex VertexId) VerticesIterable{
	iterator := func() <-chan VertexId{
		ch := make(chan VertexId)
		go func(){
			if connected,ok:=g.edges[vertex];ok{
				for vertexId,_:=range connected{
					if g.CheckEdge(vertex,vertexId){
						ch <- vertexId
					}
				}
			}
			close(ch)
		}()
		return ch
	}
	return VerticesIterable(&vertexIterableHelper{iterFunc:iterator})
}

func (g *DGraph) Reverse() *DGraph{
	r := NewDirected()

	vertices := g.VerticesIter()
	for vertex:=range vertices{
		r.AddVertex(vertex)
	}

	edges:=g.EdgesIter()
	for edge:= range edges{
		r.AddEdge(edge.To,edge.From,1)
	}

	return r
}