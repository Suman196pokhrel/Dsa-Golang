package linkedlistgo

import "fmt"

type snode struct {
	data int
	next *snode
}

type slinkedList struct {
	head *snode
}

func (s *slinkedList) length() int {
	if s.head == nil {
		return 0
	} else {
		current := s.head
		length := 0
		for current != nil {
			length++
			current = current.next
		}

		return length
	}
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

func (s *slinkedList) insertAtEnd(userData int) *slinkedList {
	node := &snode{data: userData}

	// CHECK IF SLL IS EMPTY
	if s.head == nil {
		s.head = node
	} else {
		current := s.head

		for current.next != nil {
			current = current.next
		}

		current.next = node
	}

	return s
}

func (s *slinkedList) insertAtIndex(userData int, index int) *slinkedList {
	node := &snode{data: userData}

	// CHECK IF SLL IS EMPTY
	if s.head == nil {
		s.head = node
	} else {
		// CHECK IF PROVIDED INDEX IS VALID FOR THE SLL
		fmt.Print()
	}

	return s
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
	sll.insertAtEnd(10)
	sll.insertAtEnd(9)
	sll.insertAtEnd(8)

	sll.displayElements()

	fmt.Println(sll.length())

	// INSERT AT INDEX

	// DELETE AT START

	// DELETE AT END

	// DELETE AT INDEX

	// LENGTH OF SLL

	// UPDATE AT INDEX

	// UPDATE FIRST ITEM WITH VALUE

}
