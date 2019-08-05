package Heap

import "sync"

type Item interface{
	Less(than Item) bool
}

/* 堆由一维数组表示 */
type Heap struct{
	sync.Mutex
	data []Item
	min bool
}

func New() *Heap{
	return &Heap{
		data:make([]Item,0),
	}
}

//最小堆
func NewMin() *Heap{
	return &Heap{
		data:make([Item],0),
		min:true,
	}
}

//最大堆
func NewMax() *Heap{
	return &Heap{
		data:make([]Item,0),
		min:false,
	}
}


func (h Heap) isEmpty() bool{
	return len(h.data) == 0
}

func (h Heap) Len() int{
	return len(h.data)
}

func (h Heap) Get(n int) Item{
	return h.data[n]
}

func (h *Heap) Insert(n Item){
	h.Lock()
	defer h.Unlock()

	h.data = append(h.data,n)
	//元素自动上浮，调整到堆内合适的位置
	h.siwm()

	return
}

func (h *Heap) Extract() (el Item){
	h.Lock()
	defer h.Unlock()

	if h.Len()==0{
		return
	}

	el = h.data[0]

	last := h.data[h.Len()-1]

	if h.Len()==1{
		h.data=nil
		return
	}

	h.data = append([]Item{last},h.data[1:h.Len()-1]...)
	//我们删除堆顶元素并将数组的最后一个元素放到顶端， 减小堆的大小并让这个元素下沉到合适的位置
	//将最后一个元素放到堆顶，然后下沉到合适位置
	h.sink()
	return
}

func (h *Heap) Less(a,b Item) bool{
	if h.min{
		return a.Less(b)
	}else{
		return b.Less(a)
	}
}

func (h *Heap) siwm(){
	for i,parent :=h.Len()-1,h.Len()-1;i>0;i=parent{
		parent = i>>1
		if h.Less(h.Get(i),h.Get(parent)){
			h.data[parent],h.data[i] = h.data[i],h.data[parent]
		}else{
			break
		}
	}
}

/* 0 1 2 3 4 5 6 7 8 9 10
i=0 childleft=1 childright=2
i=1 childleft=3 childright=4
i=2 childleft=5 childright=6
len=11 
*/
/* 堆：下标从0开始 */
func (h *Heap) sink() {
	for i,child :=0,1;i<h.Len() && i*2+1 <h.Len();i=child{
		child = 2*i+1

		if child+1 <= h.Len()-1&&h.Less(h.Get(child+1),h.Get(child)){
			child++
		}

		if h.Less(h.Get(i),h.Get(child)){
			break
		}

		h.data[i],h.data[child] = h.data[child],h.data[i]
	}
}


