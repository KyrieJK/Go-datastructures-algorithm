package Dijkstra

import (
	"Go-datastructures-algorithm/Graph"
	"Go-datastructures-algorithm/Priority_Queue"
)

const MAX_VALUE = 9999

func Dijkstra(g *Graph.UndGraph, source Graph.VertexId) map[Graph.VertexId]Graph.VertexId {
	visited := make(map[Graph.VertexId]bool, g.VerticesCount())
	dist := make(map[Graph.VertexId]int)
	prev := make(map[Graph.VertexId]Graph.VertexId)
	Q := Priority_Queue.NewMin() //创建最小堆，堆顶为最小元素

	vertices := g.VerticesIter()

	dist[source] = 0
	for vertex := range vertices {
		if vertex != source {
			dist[vertex] = MAX_VALUE
			prev[vertex] = 0
		}
		Q.Insert(*Priority_Queue.NewItem(vertex, dist[vertex]))
	}

	for Q.Len() > 0 {
		u := Q.Extract().Value.(Graph.VertexId) //取堆顶元素，本算法创建的为最小堆
		visited[u] = true                       //变更读取标志位

		for neighbour := range g.GetNeighbours(u).VerticesIter() {
			if !visited[neighbour] {
				//根据最近取的距离源点最近的顶点u，以u为中转点获取其邻居顶点
				//计算源点在以u为中转顶点情况下到各个neighbour顶点的距离，计算后的距离存储在alt变量中
				alt := dist[u] + g.GetEdge(u, neighbour)

				//比较alt与dist[neighbour]，如果以u为中转点计算得出的alt<dist[neighbour]，则更新dist[neighbour]
				//将neighbour顶点的前继顶点存储在prev[neighbour]数组中
				//更新最小堆中元素的优先级
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
