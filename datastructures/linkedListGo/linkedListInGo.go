package linkedlistgo

import "fmt"

type snode struct {
	data int
	next *snode
}

type slinkedList struct {
	head *snode
}

func (s *slinkedList) insertAtStart(userData int) *slinkedList {
	node := &snode{data: userData} // by default next is set to nil

	// CHECK IF LL IS EMPTY
	if s.head == nil {
		s.head = node
	} else {
		node.next = s.head
		s.head = node
	}

	return s
}

func (s *slinkedList) displayElements() {

	// CHECK IF LL IS EMPTY
	if s.head == nil {
		fmt.Println("Single Linked List is Empty")
	} else {
		current := s.head

		for current != nil {
			fmt.Print(current.data, " => ")
			current = current.next
		}
		fmt.Println()
	}
}

func SingleLinkedListInGo() {
	// INITIALIZING A SLL
	sll := slinkedList{}

	// INSERT AT START
	sll.insertAtStart(11)
	sll.insertAtStart(12)
	sll.insertAtStart(13)
	sll.insertAtStart(14)

	sll.displayElements()

	// INSERT AT END

	// INSER AT END

	// INSERT AT INDEX

	// DELETE AT START

	// DELETE AT END

	// DELETE AT INDEX

	// LENGTH OF SLL

	// UPDATE AT INDEX

	// UPDATE FIRST ITEM WITH VALUE

}
