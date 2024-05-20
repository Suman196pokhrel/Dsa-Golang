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

func (dll *dLinkedList) updateAtIndex(insertIndex int, userData int) {
	if insertIndex < 0 {
		fmt.Println("Index cannot be negative")
		return
	}

	if insertIndex == 0 {
		dll.updateStart(userData)
		return
	}

	if insertIndex >= dll.length() {
		fmt.Println("Index cannot be greater than the length of double linkedlist")
		return
	}

	current := dll.head
	for i := 0; i < insertIndex; i++ {
		current = current.next
	}

	current.data = userData
}

func (dll *dLinkedList) deleteAtIndex(nodeIndex int) {
	if dll.head == nil {
		fmt.Println("Double linkedlist is empty")
		return
	}

	// DELETE FIRST NODE
	if nodeIndex == 0 {
		dll.head = dll.head.next
		dll.head.prev = nil
		return
	}

	current := dll.head
	for i := 0; i < nodeIndex-1; i++ {
		if current.next == nil {
			fmt.Println("Index out of range")
			return
		}
		current = current.next
	}

	// current now points to the node just before the one to be deleted
	if current.next == nil {
		fmt.Println("Index out of range")
		return
	}

	toDelete := current.next
	current.next = toDelete.next
	if toDelete.next != nil {
		toDelete.next.prev = current
	}
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
	dll.updateAtIndex(6, 3434)
	dll.displayElements()

	// DELETE
	dll.deleteAtIndex(6)
	dll.displayElements()

}
