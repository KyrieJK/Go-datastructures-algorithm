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