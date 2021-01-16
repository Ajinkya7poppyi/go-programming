package main

import "fmt"

// SLLNode ...
type SLLNode struct {
	next  *SLLNode
	value int
}

// SetValue ...
func (sNode *SLLNode) SetValue(v int) {
	sNode.value = v
}

// GetValue ...
func (sNode *SLLNode) GetValue() int {
	return sNode.value
}

// NewSLLNode ...
func NewSLLNode() *SLLNode {
	return new(SLLNode)
}

// SingleLinkedList ...
type SingleLinkedList struct {
	head *SLLNode
	tail *SLLNode
}

func newSingleLinkedList() *SingleLinkedList {
	return new(SingleLinkedList)
}

//Add ...
func (list *SingleLinkedList) Add(v int) {
	newNode := &SLLNode{value: v}
	if list.head == nil {
		list.head = newNode
	} else if list.tail == list.head {
		list.head.next = newNode
	} else if list.tail != nil {
		list.tail.next = newNode
	}
	list.tail = newNode
}

//String ..
func (list *SingleLinkedList) String() string {
	s := ""
	for n := list.head; n != nil; n = n.next {
		s += fmt.Sprintf(" {%d} ", n.GetValue())
	}
	return s
}

func main() {
	list := newSingleLinkedList()
	list.Add(3)
	list.Add(5)
	list.Add(2)
	list.Add(3)
	fmt.Println("Hello, Lists", list)
}
