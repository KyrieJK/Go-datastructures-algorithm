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

func (tree *BinaryTree) SearchItem(i int) (*Node,bool){
	if tree.root == nil{
		return nil,false
	}

	currentNode := tree.root

	for currentNode != nil{
		if i == currentNode.data{
			return currentNode.data
		}else if i>currentNode.data{
			currentNode = currentNode.right
		}else if i<currentNode.data{
			currentNode = currentNode.left
		}
	}
	return nil,false
}

/* 二叉树前序遍历 */
func (tree *BinaryTree) PreorderTraversal(rootNode *Node,callback func(int)){
	if root != nil{
		callback(rootNode.data)
		PreorderTraversal(rootNode.left,callback)
		PreorderTraversal(rootNode.right,callback)
	}
}

/* 二叉树中序遍历 */
func (tree *BinaryTree) InorderTraversal(rootNode *Node,callback func(int)){
	if rootNode!=nil{
		InorderTraversal(rootNode.left,callback)
		callback(rootNode.data)
		InorderTraversal(rootNode.right,callback)
	}
}

/* 二叉树后续遍历 */
func (tree *BinaryTree) PostTraversal(rootNode *Node,callback func(int)){
	if rootNode!=nil{
		PostTraversal(rootNode.left,callback)
		PostTraversal(rootNode.right,callback)
		callback(rootNode.data)
	}
}