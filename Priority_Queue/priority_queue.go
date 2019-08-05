package Priority_Queue

import (
	"Go-datastructures-algorithm/Heap"
	"Go-datastructures-algorithm/Queue"
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

func (pq *PQ) Insert(el Item) {
	pq.data.Insert(Heap.Item(el))
}

func (pq *PQ) Extract() (el Item) {
	return pq.data.Extract().(Item) //取堆顶元素
}

func (pq *PQ) ChangePriority(val interface{}, priority int) {
	var storage = Queue.New()

	popped := pq.Extract()
	for val != popped.Value {
		if pq.Len() == 0 {
			panic("Item not found")
		}

		storage.Push(popped)
		popped = pq.Extract()
	}

	popped.Priority = priority
	pq.data.Insert(popped)

	for storage.Len() > 0 {
		pq.data.Insert(storage.ExtractFirst().(Heap.Item))
	}
}
