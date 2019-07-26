package main

import "fmt"

type Node struct {
	isKey    bool
	parent   *Node
	children map[int32]*Node
}

type Trie struct {
	root *Node
}

func (trie *Trie) Init() {
	trie.root = &Node{children:map[int32]*Node{}}
}

func (trie Trie) Add(item string){
	currentNode := trie.root
	for _,r := range item{
		if _,ok := currentNode.children[r];ok{
			currentNode = currentNode.children[r]
		}else{
			node := &Node{children:map[int32]*Node{},parent:currentNode}
			currentNode.children[r] = node
			currentNode = node
		}
	}
	currentNode.isKey = true
}

func (trie Trie) Search(item string) bool{
	currentNode := trie.root
	for _,r:=range item{
		if _,ok := currentNode.children[r];ok{
			currentNode = currentNode.children[r]
		}else {
			return false
		}
	}

	if currentNode.isKey == false{
		return false
	}
	return true
}

func (trie Trie) Remove(item string) bool{
	currentNode :=trie.root
	for _,r := range item{
		if _,ok := currentNode.children[r];ok{
			currentNode = currentNode.children[r]
		}else{
			return false
		}
	}

	if currentNode.isKey == false{
		return false
	}

	currentNode.isKey = false
	symbolIndex := len(item) - 1
	for len(currentNode.children) == 0{
		currentNode = currentNode.parent
		delete(currentNode.children,int32(item[symbolIndex]))
		symbolIndex--
	}
	return true
}

func main(){
	item := "hello world"
	for index,value := range item{
		fmt.Println(index,value)
	}
}
