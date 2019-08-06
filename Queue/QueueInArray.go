package Queue

import "sync"

type QueueInArray struct {
	queue []interface{}
	len   int
	lock  *sync.Mutex
}

func New() *QueueInArray {
	queue := &QueueInArray{}
	queue.queue = make([]interface{}, 0)
	queue.len = 0
	queue.lock = new(sync.Mutex)

	return queue
}

func (q QueueInArray) Len() int {
	return q.len
}

func (q QueueInArray) isEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.len == 0
}

func (q *QueueInArray) ExtractFirst() (el interface{}) {
	el = q.queue[0]
	q.queue = q.queue[1:]
	q.len--
	return el
}

func (q *QueueInArray) Push(el interface{}) {
	q.queue = append(q.queue, el)
	q.len++

	return
}

func (q *QueueInArray) Peek() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.queue[0]
}
