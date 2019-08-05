package Queue

import "sync"

type Queue struct {
	queue []interface{}
	len   int
	lock  *sync.Mutex
}

func New() *Queue {
	queue := &Queue{}
	queue.queue = make([]interface{}, 0)
	queue.len = 0
	queue.lock = new(sync.Mutex)

	return queue
}

func (q Queue) Len() int {
	return q.len
}

func (q Queue) isEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.len == 0
}

func (q *Queue) ExtractFirst() (el interface{}) {
	el = q.queue[0]
	q.queue = q.queue[1:]
	q.len--
	return el
}

func (q *Queue) Push(el interface{}) {
	q.queue = append(q.queue, el)
	q.len++

	return
}

func (q *Queue) Peek() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.queue[0]
}
