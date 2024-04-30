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

	// CHECK IF PROVIDED INDEX IS VALID FOR THE SLL
	if index < 0 {
		fmt.Println("Invalid index provided")
		return s
	}

	// CHECK IF SLL IS EMPTY OR INSERTING AT THE BEGINNING
	if s.head == nil || index == 0 {
		s.insertAtStart(userData)
		return s
	}

	current := s.head
	for i := 0; i < index-1 && current != nil; i++ {
		current = current.next
	}

	if current == nil {
		fmt.Println("Invalid index provided")
		return s
	}

	node.next = current.next
	current.next = node

	return s
}

func (s *slinkedList) deleteAtStart() *slinkedList {

	if s.head == nil {
		fmt.Println("SLL is Empty")
		return s
	}

	s.head = s.head.next

	return s
}

func (s *slinkedList) deleteAtEnd() *slinkedList {

	if s.head == nil {
		fmt.Println("SLL is empty")
		return s
	}

	// If the list has only one node
	if s.head.next == nil {
		s.head = nil
		return s
	}

	// Find the last node
	current := s.head
	for current.next.next != nil {
		current = current.next
	}

	// Delete the last node
	current.next = nil

	return s
}

func (s *slinkedList) deleteAtIndex(index int) *slinkedList {

	// CHECK IF SLL IS EMPTY
	if s.head == nil {
		fmt.Println("SLL is empty")
		return s
	}

	if index == 0 {
		s.head = s.head.next
		return s
	}

	current := s.head
	for i := 0; i < index-1 && current != nil; i++ {
		current = current.next
	}

	if current == nil || current.next == nil {
		fmt.Println("Invalid index")
		return s
	}

	current.next = current.next.next

	return s
}

func (s *slinkedList) updateAtStart() *slinkedList {
	return s
}

func (s *slinkedList) updateAtEnd() *slinkedList {
	return s
}

func (s *slinkedList) updateAtIndex() *slinkedList {
	return s
}

func (s *slinkedList) updateValue() *slinkedList {
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
	sll.insertAtIndex(7, 8)
	sll.displayElements()

	// DELETE AT START
	sll.deleteAtStart()
	sll.displayElements()

	// DELETE AT END
	sll.deleteAtEnd()
	sll.displayElements()

	// DELETE AT INDEX
	sll.deleteAtIndex(2)
	sll.displayElements()

	// LENGTH OF SLL
	sll.length()

	// UPDATE AT INDEX

	// UPDATE FIRST ITEM WITH VALUE

}
