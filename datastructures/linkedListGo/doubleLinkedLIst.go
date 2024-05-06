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

func DoubleLinkedListInGo() {
	dll := dLinkedList{}

	// INSERT AT START
	dll.insertAtBeginning(11)
	dll.insertAtBeginning(12)
	dll.insertAtBeginning(13)
	dll.insertAtBeginning(14)
	dll.insertAtBeginning(15)
	dll.displayElements()

}
