package trees

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) Insert(newValue int) {
	if newValue <= n.value {
		// IF NEW VALUE IS SMALLER MOVE LEFT
		if n.left == nil {
			n.left = &Node{value: newValue}
		} else {
			n.left.Insert(newValue)
		}

	} else if newValue >= n.value {
		// IF NEW VLAUE LARGER MIVR RIGHT
		if n.right == nil {
			n.right = &Node{value: newValue}
		} else {
			n.right.Insert(newValue)
		}
	}
}

func (n *Node) Search(value int) bool {
	if n == nil {
		return false
	}

	if value < n.value {
		// MOVE LEFT
		return n.left.Search(value)
	} else if value > n.value {
		// MOVE RIGHT
		return n.right.Search(value)
	}

	return true // Value found at current node
}

func BSTDemo() {

	// INITIALIZE A NODE (BECOMES THE ROOT NOE)
	tree := &Node{value: 100}

	// INSERT
	tree.Insert(50)
	tree.Insert(200)
	tree.Insert(300)
	tree.Insert(100)
	tree.Insert(10)
	tree.Insert(20)

	tree.Insert(400)

	fmt.Println(tree)

	// SEARCH
	fmt.Println(tree.Search(22))

}
