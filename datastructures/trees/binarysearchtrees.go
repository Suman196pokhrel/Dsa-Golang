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

func BSTDemo() {
	tree := &Node{value: 100}
	tree.Insert(50)
	tree.Insert(200)
	tree.Insert(300)

	fmt.Println(tree)

}
