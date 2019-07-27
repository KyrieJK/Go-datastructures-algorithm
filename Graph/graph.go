package main

import "fmt"

type VertexId uint

type Vertices []VertexId

type Edge struct{
	From VertexId
	To VertexId
}

type graph struct{
	edges map[VertexId]map[VertexId]uint
	edgesCount uint
	isDirected bool
}



func main(){
	m1 := make(map[VertexId]map[VertexId]int)
	m2 := make(map[VertexId]int)
	m2[1]=2
	m1[1]=m2
	m1[2] = m2
	for k,v := range m1{
		fmt.Println(k,v)
	}
}