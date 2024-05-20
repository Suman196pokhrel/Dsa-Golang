package linkedlistgo

import (
	"errors"
	"fmt"
)

type dnode struct {
	next *dnode
	prev *dnode
	data int
}

type dLinkedList struct {
	head *dnode
}

func (dll *dLinkedList) displayElements() error {
	if dll.head == nil {
		return errors.New("double linked list is empty")
	}

	current := dll.head
	for current != nil {
		fmt.Print(current.data, " => ")
		current = current.next
	}
	fmt.Println()

	return nil
}

func (dll *dLinkedList) length() int {
	if dll.head == nil {
		return 0
	}

	curr := dll.head
	res := 0
	for curr != nil {
		res++
		curr = curr.next
	}

	return res
}

func (dll *dLinkedList) insertAtBeginning(userData int) {
	node := &dnode{
		data: userData,
	}

	if dll.head == nil {
		fmt.Println("Double linked list is empty")
		dll.head = node
		return
	}

	node.next = dll.head
	dll.head = node
}

func (dll *dLinkedList) insertAtEnd(userData int) {

	if dll.head == nil {
		fmt.Println("Double LinkedList is empty ! , inserting at start")
		dll.insertAtBeginning(userData)
		return
	} else {
		node := dnode{data: userData}
		current := dll.head

		for current.next != nil {
			current = current.next
		}

		node.prev = current
		current.next = &node

		return
	}
}

func (dll *dLinkedList) insertInIndex(insertIndex int, userData int) {

	if insertIndex < 0 {
		fmt.Println("Index cannot be negative")
		return
	}

	if dll.head == nil || insertIndex == 0 {
		fmt.Println("Double LinkedList is empty ! inserting at first")
		dll.insertAtBeginning(userData)
		return
	}

	if insertIndex >= dll.length() {
		fmt.Println("Given index is out of range , inserting at end")
		dll.insertAtEnd(userData)
		return
	}

	node := dnode{
		data: userData,
	}
	current := dll.head

	for i := 0; i < insertIndex-1; i++ {
		current = current.next
	}

	// NODE BEING INSERTED
	node.next = current.next
	node.prev = current

	// AFTER NODE
	if current.next != nil {
		current.next.prev = &node
	}
	// BEFORE NODE
	current.next = &node
}

func (dll *dLinkedList) updateStart(userData int) {
	if dll.head == nil {
		fmt.Println("Double Linked List is empty!")
		return
	}

	dll.head.data = userData
}

func (dll *dLinkedList) updateEnd(userData int) {
	if dll.head == nil {
		fmt.Println("Double Linked List is empty!")
		return
	}

	current := dll.head
	for current.next != nil {
		current = current.next
	}
	current.data = userData

}

func DoubleLinkedListInGo() {
	dll := dLinkedList{}

	// INSERTS
	for i := 0; i < 5; i++ {
		dll.insertAtBeginning(i)
	}
	dll.displayElements()
	dll.insertAtEnd(10)
	dll.displayElements()
	dll.insertInIndex(5, 222)
	dll.displayElements()

	// UPDATE
	dll.updateStart(1919)
	dll.updateEnd(10000)
	dll.displayElements()

}
