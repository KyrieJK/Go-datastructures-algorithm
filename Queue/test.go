package main

import "fmt"
import "runtime"


type People struct{}

func (p *People) ShowA(){
	fmt.Println("showA")
	p.ShowB()
}

func (p *People) ShowB(){
	fmt.Println("showB")
}

type Teacher struct{
	People
}

func (t *Teacher) ShowB(){
	fmt.Println("teacher showB")
}

func main(){
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int,1)
	string_chan :=make(chan string,1)
	int_chan<-1
	string_chan<-"hello"
	select{
	case value:=<-int_chan:
		fmt.Println(value)
	case value:=<-string_chan:
		fmt.Println(value)
	}
}

// type student struct{
// 	Name string
// 	Age int
// }

// func pase_student(){
// 	m:=make(map[string]*student)

// 	stus:=[]student{
// 		{Name:"zhou",Age:24},
// 		{Name:"li",Age:23},
// 		{Name:"wang",Age:22},
// 	}

// 	for _,stu:=range stus{
// 		m[stu.Name] = &stu
// 		fmt.Printf("%p\n",m[stu.Name])
// 	}
	
// }
