package rb_tree

import (
	"fmt"
	"testing"
)

func TestRBTree(t *testing.T) {
	RBTree := NewRBTree()

	RBTree.Insert(1, 10)
	RBTree.Insert(2, 20)
	RBTree.Insert(3, 30)
	RBTree.Insert(4, 40)
	RBTree.Insert(5, 50)

	node := RBTree.Search(3)
	if node != nil {
		fmt.Println("Found key 3 with value:", node.Value)
	} else {
		fmt.Println("Key 3 not found")
	}

	RBTree.Delete(3)
	node = RBTree.Search(3)
	if node != nil {
		fmt.Println("Found key 3 with value:", node.Value)
	} else {
		fmt.Println("Key 3 not found")
	}
}
