package binarysearch

import (
	"errors"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BST struct {
	Root *Node
}

func (bst *BST) Create(value int) *BST {
	node := Node{Value: value}

	if bst.Root == nil {
		bst.Root = &node
	} else {
		bst.Root.Insert(node)
	}

	return bst
}

func (bst *BST) Find(value int) (int, error) {
	if bst.Root == nil {
		return -1, errors.New("no root")
	} else {
		node := bst.Root.FindNode(value)

		if node != nil {
			return node.Value, nil
		}

		return -1, errors.New("not found node in this tree")
	}
}

func (bst *BST) FindMinNode() *Node {
	if bst.Root == nil {
		return nil
	}

	currentNode := bst.Root

	for currentNode.Left != nil {
		currentNode = currentNode.Left
	}

	return currentNode
}

func (bst *BST) FindMaxNode() *Node {
	if bst.Root == nil {
		return nil
	}

	currentNode := bst.Root

	for currentNode.Right != nil {
		currentNode = currentNode.Right
	}

	return currentNode
}

func (currentNode *Node) Insert(newNode Node) {
	if newNode.Value < currentNode.Value {
		if currentNode.Left == nil {
			currentNode.Left = &newNode
		} else {
			currentNode.Left.Insert(newNode)
		}
	} else {
		if currentNode.Right == nil {
			currentNode.Right = &newNode
		} else {
			currentNode.Right.Insert(newNode)
		}
	}
}

func (currentNode *Node) FindNode(value int) *Node {
	if currentNode.Value == value {
		return currentNode
	} else if value < currentNode.Value && currentNode.Left != nil {
		return currentNode.Left.FindNode(value)
	} else if value > currentNode.Value && currentNode.Right != nil {
		return currentNode.Right.FindNode(value)
	}

	return nil
}
