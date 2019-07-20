package Stack

type Node struct{
	data int
	next *Node
}

type Stack struct{
	top *Node
}

func (s *Stack) Push(i int){
	data := &Node{data:i}
	if s.top !=nil{
		data.next = s.top
	}
	s.top = data
}

func (s *Stack) Pop() (int,bool){
	if s.top==nil{
		return -1,false
	}
	i := s.top.data
	s.top = s.top.next
	return i, true
}

func (s *Stack) Peek() (int,bool){
	if s.top == nil{
		return -1,false
	}

	return s.top.data,true
}

func (s *Stack) Get() []int{
	var items []int
	current := s.top
	for current !=nil{
		items = append(items,current.data)
		current = current.next
	}
	return items
}

func (s *Stack) IsEmpty() bool{
	return s.top == nil
}

func (s *Stack) Empty(){
	s.top = nil
}