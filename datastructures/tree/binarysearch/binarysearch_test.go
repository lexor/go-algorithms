package binarysearch

import "testing"

func TestBinarySearchTreeCreate(t *testing.T) {
	bst := BST{}

	if bst.Root != nil {
		t.Errorf("root defined")
	}

	// Let's add node
	bst.Create(5)

	if bst.Root.Left != nil {
		t.Errorf("root node's left node is defined")
	}

	if bst.Root.Right != nil {
		t.Errorf("root node's right node is defined")
	}
}

func TestBinarySearchTreeInsert(t *testing.T) {
	node := Node{Value: 10}
	node2 := Node{Value: 20}

	node.Insert(node2)

	if node.Right == nil {
		t.Errorf("right node is nil")
	}

	if node.Left != nil {
		t.Errorf("left node is not nil")
	}
}

func TestBinarySearchTreeFind(t *testing.T) {
	bst := BST{}

	bst.Create(10).Create(20).Create(5)

	_, err := bst.Find(20)
	if err != nil {
		t.Errorf("node not found")
	}

	_, err = bst.Find(10)
	if err != nil {
		t.Errorf("node not found")
	}

	_, err = bst.Find(5)
	if err != nil {
		t.Errorf("node not found")
	}
}
