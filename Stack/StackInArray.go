package Stack

import "sync"

type Stack struct {
	stack []interface{}
	len int
	lock sync.Mutex
}

func New() *Stack{
	s := &Stack{}
	s.stack = make([]interface{},0)
	s.len = 0

	return s
}

func (s Stack) Len() int{
	s.lock.Lock()

	defer s.lock.Unlock()

	return s.len
}

func (s Stack) isEmpty() bool{
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.len==0
}

func (s *Stack) Pop() (ele interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	ele,s.stack = s.stack[0],s.stack[1:]
	s.len--
	return ele
}

func (s *Stack) Push(ele interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	prepend := make([]interface{},1)
	prepend[0] = ele
	s.stack = append(prepend,s.stack...)
	s.len++
}

func (s Stack) Peek() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.stack[0]
}
