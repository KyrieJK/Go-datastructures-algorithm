package Graph

/**
无向图类
*/

type UndGraph struct{
	graph
}

/* 无向图实例 */
func NewUndirected() *UndGraph{
	return &UndGraph{
		graph{
			edgesCount:0,
			edges:make(map[VertexId]map[VertexId]int),
			isDirected:false,
		},
	}
}

