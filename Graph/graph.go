package main

import (
	"errors"
	"fmt"
)

type VertexId uint

type Vertices []VertexId

type Edge struct {
	From VertexId
	To   VertexId
}

type graph struct {
	edges      map[VertexId]map[VertexId]uint //graph中边的表示数据结构
	edgesCount uint
	isDirected bool //标识有向、无向图
}

/* 所有边的遍历 */
func (g *graph) EdgesIter() <-chan Edge {
	ch := make(chan Edge)
	go func() {
		for from, connectedVertices := range g.edges {
			for to, _ := range connectedVertices {
				if g.isDirected {
					ch <- Edge{from, to}
				} else {
					if from < to {
						ch <- Edge{from, to}
					}
				}
			}
		}
		close(ch)
	}()
	return ch
}

func (g *graph) VerticesIter() <-chan VertexId {
	ch := make(chan VertexId)
	go func() {
		for vertex, _ := range g.edges {
			ch <- vertex
		}
		close(ch)
	}()
	return ch
}

func (g *graph) CheckVertex(vertex VertexId) bool {
	_, exists := g.edges[vertex]
	return exists
}

func (g *graph) TouchVertex(vertex VertexId) {
	if _, ok := g.edges[vertex]; !ok {
		g.edges[vertex] = make(map[VertexId]uint)
	}
}

func (g *graph) AddVertex(vertex VertexId) error {
	i, _ := g.edges[vertex]
	if i != nil {
		return errors.New("Vertex already exists")
	}

	g.edges[vertex] = make(map[VertexId]uint)
	return nil
}

func (g *graph) RemoveVertex(vertex VertexId) error {
	if !g.CheckVertex(vertex) {
		return errors.New("Unkknown vertex")
	}

	delete(g.edges, vertex)

	for _, connectedVertices := range g.edges {
		delete(connectedVertices, vertex)
	}
	return nil
}

func (g *graph) VerticesCount() int {
	return len(g.edges)
}

func (g *graph) AddEdge(from, to VertexId, weight int) error {
	if from == to{
		return errors.New("Can't add self loop")
	}

	if !g.CheckVertex(from) || !g.CheckVertex(to){
		return errors.New("Vertices don't exist")
	}

	i,_:=g.edges[from][to]
	j,_:=g.edges[to][from]

	if i>0||j>0{
		return errors.New("Edge already exists")
	}

	g.TouchVertex(from)
	g.TouchVertex(to)

	g.edges[from][to] = weight

	if !g.isDirected{
		g.edges[to][from] = weight
	}

	g.edgesCount++

	return nil
}

func main() {
	m1 := make(map[VertexId]map[VertexId]int)
	m2 := make(map[VertexId]int)
	m2[1] = 2
	m2[2] = 3
	m1[1] = m2
	i,_:=m1[1][2]
	fmt.Println(i)
}
