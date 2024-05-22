package trees

import (
	"fmt"
)

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

func (n *Node) InOrderTraversal() {
	if n != nil {
		n.left.InOrderTraversal()
		fmt.Println(n.value)
		n.right.InOrderTraversal()
	}

}

func (n *Node) PreOrderTraversal() {
	if n != nil {
		fmt.Println(n.value)
		n.left.InOrderTraversal()
		n.right.InOrderTraversal()
	}

}

func (n *Node) PostOrderTraversal() {
	if n != nil {
		n.left.InOrderTraversal()
		n.right.InOrderTraversal()
		fmt.Println(n.value)

	}

}

// PrintBSTWithLevel prints the BST with each node's value and its level
func PrintBSTWithLevel(root *Node, level int) {
	if root == nil {
		return
	}

	PrintBSTWithLevel(root.left, level+1)
	fmt.Printf("Level %d: %d\n", level, root.value)
	PrintBSTWithLevel(root.right, level+1)
}

func findSmallest(n *Node) *Node {
	current := n

	if current.left != nil {
		current = current.left
	}

	return current
}

func (n *Node) DeleteNode(value int) *Node {

	if n == nil {
		return nil
	}

	// Step 1 :  Find the node to be deleted
	if value < n.value {
		n.left = n.left.DeleteNode(value)
	} else if value > n.value {
		n.right = n.right.DeleteNode(value)
	} else {
		// NODE TO BE DELETED FOUND

		// CASE 1 : Node is a leaf node
		if n.left == nil && n.right == nil {
			return nil
		}

		// CASE 2 : Node has one child
		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		}

		// CASE 3 : Node has two children
		// Find the inorder successor (smallest in the right subtreee)
		minRight := findSmallest(n.right)
		// replace the mod's value with the successor's value
		n.value = minRight.value
		// DEELTE the successor
		n.right = n.right.DeleteNode(minRight.value)

	}

	return n

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

	// PrintBSTWithLevel(tree, 0)

	// NOTE : Inorder traversal always gives you a sorted sequence of values
	// tree.InOrderTraversal()

	tree.PostOrderTraversal()
}
