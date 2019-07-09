package Queue

type Queue struct{
	element *Node
}

type Node struct{
	data int
	next *Node
}

/* 获取队列中所有的元素 */
func (list *Queue) Get() []int{
	var items []int
	current :=list.element
	for current!= nil{
		items = append(items,current.data)
		current = current.next
	}
	return items
}

func (list *Queue) Peek() (int,bool){
	if list.element ==nil{
		return 0,false
	}
	return list.element.data,true
}

func (list *Queue) IsEmpty() bool{
	return list.element == nil
}

func (list *Queue) Enqueue(i int){
	data := &Node{data:i}
	if list.element != nil{
		data.next = list.element
	}
	list.element = data
}

func (list *Queue) Dequeue() (int,bool){
	if list.element == nil{
		return 0,false
	}

	if list.element.next == nil{
		i := list.element.data
		list.element = nil
		return i,true
	}

	current := list.element
	/* 队列先入先出，这里是寻找队列最先进入的那一个元素 */
	for{
		if current.next.next == nil{
			i := current.next.data
			current.next = nil
			return i,true
		}
		current = current.next
	}
}