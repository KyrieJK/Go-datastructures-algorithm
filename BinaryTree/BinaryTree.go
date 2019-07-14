package BinaryTree

type Node struct {
	data   int
	parent *Node
	left   *Node
	right  *Node
}

type BinaryTree struct {
	root *Node
}

func (tree *BinaryTree) InsertItem(i int) {
	if tree.root == nil {
		tree.root = &Node{data: i}
		return
	}

	currentNode := tree.root
	for {
		if i > currentNode.data {
			currentNode.right == nil{
				currentNode.right = &Node{data:i,parent:currentNode}
				return
			}
			currentNode = currentNode.right
		}else{
			if currentNode.left == nil{
				currentNode.left = &Node{data:i,parent:currentNode}
				return
			}
			currentNode = currentNode.left
		}
	}
}
