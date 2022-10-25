package pointers

import (
	"errors"
	"fmt"
)

type LinkedList interface {
	Push(data interface{}) error
	Get(index int) (*Node, error)
	Insert(data interface{}, position int) error
	Remove(index int) (*Node, error)
	Search(data interface{}) (*Node, error)
	Size() (int, error)
}

type linkedList struct {
	head *Node
}

func NewLinkedList() LinkedList {
	return &linkedList{}
}

func (ll *linkedList) Push(data interface{}) (err error) {
	err = nil
	newNode := NewNode(data)
	newNode.NextNode = ll.head
	ll.head = newNode
	return
}

func (ll *linkedList) Size() (size int, err error) {
	err = nil
	current := ll.head
	size = 0

	for current != nil {
		size++
		current = current.NextNode
	}

	return
}

func (ll *linkedList) Get(index int) (node *Node, err error) {
	err = nil
	node = ll.head
	position := 0

	for index != position {
		position++
		node = node.NextNode
	}

	return
}
func (ll *linkedList) Remove(index int) (*Node, error) {
	if ll.head == nil {
		return nil, errors.New("the list is empty")
	}

	if index == 0 {
		removed := ll.head
		ll.head = ll.head.NextNode
		return removed, nil
	}

	current := ll.head

	for index > 1 {
		index--
		current = current.NextNode
	}

	prev := current
	next := current.NextNode
	prev.NextNode = next.NextNode

	return next, nil
}

func (ll *linkedList) Insert(data interface{}, index int) (err error) {
	err = nil
	current := ll.head

	if index == 0 {
		ll.Push(data)
		return
	}

	for index > 1 {
		index--
		if current.NextNode == nil {
			return fmt.Errorf("index %v out of bound", index)
		}
		current = current.NextNode
	}

	prev := current
	next := current.NextNode
	new := NewNode(data)
	prev.NextNode = new
	new.NextNode = next

	return
}

func (ll *linkedList) Search(data interface{}) (*Node, error) {
	current := ll.head

	for current != nil {
		if current.Data == data {
			return current, nil
		}
		current = current.NextNode
	}

	return nil, fmt.Errorf("%s not found", data)
}
