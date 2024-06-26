package linkedlistgo

import (
	"errors"
	"fmt"
)

type SortMode string

const (
	Ascending  SortMode = "ASC"
	Descending SortMode = "DES"
)

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

// func (s *slinkedList) String() string {
// 	// CHECK IF LL IS EMPTY
// 	if s.head == nil {
// 		fmt.Println("Single Linked List is Empty")
// 		return "Single Linked List is Empty"

// 	} else {
// 		current := s.head
// 		var output strings.Builder // Using strings.Builder for efficient string concatenation
// 		for current != nil {
// 			fmt.Print(current.data, " => ")
// 			output.WriteString(fmt.Sprintf("%v => ", current.data))
// 			current = current.next
// 		}
// 		return output.String()
// 	}
// }

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

func (s *slinkedList) updateAtStart(userData int) *slinkedList {

	// CHECK IF SLL IS EMPTY
	if s.head == nil {
		fmt.Println("SLL IS EMPTY , CANNOT UPDATE")
		return s
	}

	s.head.data = userData

	return s
}

func (s *slinkedList) updateAtEnd(userData int) *slinkedList {
	// CHECK IF SLL IS EMPTY
	if s.head == nil {
		fmt.Println("SLL IS EMPTY , CANNOT UPDATE")
		return s
	}

	current := s.head
	for current.next != nil {
		current = current.next
	}

	current.data = userData
	return s
}

func (s *slinkedList) updateAtIndex(index int, userData int) *slinkedList {

	// CHECK IF SLL IS EMPTY
	if s.head == nil {
		fmt.Println("SLL IS EMPTY , CANNOT UPDATE")
		return s
	}

	if index == 0 {
		s.head.data = userData
		return s
	}

	if index >= s.length() {
		fmt.Println("Invalid index , CANNOT UPDATE")
		return s
	}

	current := s.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	current.data = userData

	return s
}

func (s *slinkedList) updateValue(targetElement int, userData int) *slinkedList {
	if s.head == nil {
		fmt.Print("SLL IS EMPTY , CANNOT UPDATE")
		return s
	}

	current := s.head
	for current != nil {
		if current.data == targetElement {
			current.data = userData
			return s
		}

		current = current.next
	}

	fmt.Println("ELEMENT WAS NOT FOUND IN SLL")

	return s
}

func (sm SortMode) IsValid() bool {
	return sm == Ascending || sm == Descending
}

func (s *slinkedList) swapAdjacentNodes(indexOfFirstNode int) error {
	if indexOfFirstNode < 0 || indexOfFirstNode >= s.length()-1 {
		return errors.New("index out of range")
	}

	current := s.head
	for i := 0; i < indexOfFirstNode; i++ {
		current = current.next
	}

	current.data, current.next.data = current.next.data, current.data
	return nil
}

func (s *slinkedList) sort(mode SortMode) *slinkedList {

	if !mode.IsValid() {
		fmt.Println("INVALID SORTING MODE")
		return s
	}

	if s.head == nil {
		fmt.Println("SLL IS EMPTY")
		return s
	}

	// TODO :BUBBLE SORT IMPLEMENTATION
	var sorted bool
	var lastSwapped *snode = nil
	for !sorted {
		sorted = true
		prev := s.head
		current := s.head.next

		for current != lastSwapped {
			if prev.data > current.data {
				// Swap data of adjacent nodes
				prev.data, current.data = current.data, prev.data
				sorted = false // Set sorted to false to indicate that a swap occurred
			}
			prev = current
			current = current.next
		}
		lastSwapped = prev // Update lastSwapped to optimize next pass
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
	sll.insertAtIndex(7, 8)
	sll.displayElements()

	// DELETE AT START
	// sll.deleteAtStart()
	sll.displayElements()

	// DELETE AT END
	// sll.deleteAtEnd()
	sll.displayElements()

	// DELETE AT INDEX
	// sll.deleteAtIndex(2)
	sll.displayElements()

	// LENGTH OF SLL
	fmt.Println(sll.length())
	// UPDATE FIRST
	sll.updateAtStart(100)
	sll.displayElements()

	// UPDATE LAST
	sll.updateAtEnd(999)
	sll.displayElements()

	// UPDATE AT INDEX
	sll.updateAtIndex(2, 222)
	sll.displayElements()

	// UPDATE FIRST ITEM WITH VALUE
	sll.updateValue(99999, 2020)
	sll.displayElements()

	// SORTING SLL
	sll.sort(Ascending)
	sll.displayElements()

}
