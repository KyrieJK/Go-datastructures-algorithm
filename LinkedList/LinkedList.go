package LinkedList

/* 单向链表数据结构 */

type Node struct{
	data int
	next *Node
}

type LinkedList struct{
	head *Node
}

/* 链表头部插入元素 */
func (list *LinkedList) InsertFirst(i int){
	data := &Node{data:i}
	if list.head != nil{
		data.next = list.head
	}
	list.head = data
}

/* 链表尾部插入元素 */
func (list *LinkedList) InsertLast(i int){
	data := &Node{data:i}
	if list.head == nil{
		list.head = data
		return
	}
	current := list.head
	for current.next != nil{
		current = current.next
	}
	current.next = data
}

/* 根据value值移除链表中的元素 */
func (list *LinkedList) RemoveByValue(i int) bool{
	if list.head == nil{
		return false
	}
	if list.head.data == i{
		list.head = list.head.next
		return true
	}
	current := list.head
	for current.next != nil{
		if current.next.data == i{
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}

/* 通过下标索引删除链表中的元素 */
func (list *LinkedList) RemoveByIndex(i int) bool{
	if list.head == nil{
		return false
	}

	if i<0{
		return false
	}
	if i==0{
		list.head = list.head.next
		return true
	}
	current := list.head
	for j:=1;j<i;j++{
		if current.next.next == nil{
			return false
		}

		current = current.next
	}
	current.next = current.next.next
	return true
}

func (list *LinkedList) SearchValue(i int) bool{
	if list.head == nil{
		return false
	}

	current := list.head
	for current != nil{
		if current.data == i{
			return true
		}
		current = current.next
	}

	return false
}

func (list *LinkedList) GetFirst() (int,bool){
	if list.head == nil{
		return 0,false
	}
	return list.head.data,true
}

func (list *LinkedList) GetLast() (int,bool){
	if list.head == nil{
		return 0, false
	}

	current := list.head
	for current.next != nil{
		current = current.next
	}
	return current.data,true
}

func (list *LinkedList) GetSize() int{
	count := 0
	current := list.head
	for current != nil{
		current = current.next
		count++
	}
	return count
}

func (list *LinkedList) GetItems() []int{
	var items []int
	current := list.head
	for current != nil{
		items = append(items,current.data)
		current = current.next
	}
	return items
}