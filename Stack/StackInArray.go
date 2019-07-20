package Stack

const arraySize = 10

type Stack struct {
	topIndex  int
	data [arraySize]int
}

/* 入栈 */
func (s *Stack) Push(i int) bool {
	if s.topIndex == len(s.data){
		return false
	}
	s.data[s.topIndex] = i
	s.topIndex +=1
	return true
}

/* 出栈 */
func (s *Stack) Pop() (int,bool){
	if s.topIndex == 0{
		return 0,false
	}
	i := s.data[s.topIndex-1]
	s.topIndex -=1
	return i,true
}

/* 取顶端元素 */
func (s *Stack) Peek() int{
	return s.data[s.topIndex-1]
}

/* 获取栈中存在的所有元素 */
func (s *Stack) Get() []int{
	return s.data[:s.topIndex]
}

/* 判断栈是否为空 */
func (s *Stack) IsEmpty() bool{
	return s.topIndex==0
}

/* 重新置栈顶索引为空，重新向栈中添加元素 */
func (s *Stack) Empty() {
	s.topIndex = 0
}
