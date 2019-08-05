package Priority_Queue

import (
	"Go-datastructures-algorithm/Heap"
)

type Item struct {
	Value    interface{}
	Priority int
}

func NewItem(value interface{}, priority int) (i *Item) {
	return &Item{
		Value:    value,
		Priority: priority,
	}
}

func (x Item) Less(than Heap.Item) bool {
	return x.Priority < than.(Item).Priority
}

type PQ struct {
	data Heap.Heap
}

func NewMax() (q *PQ) {
	return &PQ{
		data: *Heap.NewMax(),
	}
}

func NewMin() (q *PQ) {
	return &PQ{
		data: *Heap.NewMin(),
	}
}

func (pq *PQ) Len() int {
	return pq.data.Len()
}
