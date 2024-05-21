package trees

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) Insert(newValue int) {
	if newValue <= n.value {
		// MOVE TO LEFT NODE
		if n.left == nil {
			n.left = &Node{value: newValue}
			return
		} else {
			n.left.Insert(newValue)
		}

	} else if newValue >= n.value {
		// MOVE TO RIGHT NODE
		if n.right == nil {
			n.right = &Node{value: newValue}
			return
		} else {
			n.right.Insert(newValue)
		}
	}
}

func (n *Node) SearchRecursive(searchVal int) bool {

	if searchVal < n.value {
		// MOVE TO LEFT NODE
		if n.left == nil {
			return false
		} else {
			return n.left.SearchRecursive(searchVal)
		}
	} else if searchVal > n.value {
		// MOVE TO RIGHT NODE
		if n.right == nil {
			return false
		} else {
			return n.right.SearchRecursive(searchVal)
		}
	}

	return true
}

func (n *Node) SearchNonRecursive(searchVal int) bool {
	current := n

	for current != nil {
		if searchVal < current.value {
			current = current.left
		} else if searchVal > current.value {
			current = current.right
		} else {
			return true
		}
	}

	return false
}

func BSTDemo() {
	tree := &Node{
		value: 500,
	}

	// INSERT
	tree.Insert(250)
	tree.Insert(750)

	// Depth 2
	tree.Insert(125)
	tree.Insert(375)
	tree.Insert(625)
	tree.Insert(875)

	// Depth 3
	tree.Insert(62)
	tree.Insert(187)
	tree.Insert(312)
	tree.Insert(437)
	tree.Insert(562)
	tree.Insert(687)
	tree.Insert(812)
	tree.Insert(937)

	// SEARCH
	fmt.Println(tree.SearchRecursive(562222))
	fmt.Println(tree.SearchNonRecursive(562222)) // PREVENTS STACK OVERFLOW

}
