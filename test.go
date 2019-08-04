package main

import (
	"fmt"
)

func main() {
	s:=[]byte("")

	s1:=append(s,'a')
	s2:=append(s,'b')

	//fmt.Println(cap(s),len(s),string(s),s)
	fmt.Println(string(s1),string(s2),&s1[0],&s2[0])
}
