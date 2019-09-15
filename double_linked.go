package doublelinked

import (
	"errors"
	"fmt"
)

const (
	errorIndex = "Error invalide index"
)

type node struct {
	Value int
	Next  *node
	Prev  *node
}

type DoubleLinkedList struct {
	Head   *node
	Tail   *node
	Length int
}

func (d *DoubleLinkedList) indexExist(index int) bool {
	if index < 0 || index >= d.Length {
		return false
	}
	return true
}

func (d *DoubleLinkedList) traverse(index int) *node {
	n := d.Head
	for i := 0; i <= index; i++ {
		if index == i {
			return n
		}
		n = n.Next
	}
	return nil
}

func (d *DoubleLinkedList) Insert(value int) {

	n := &node{Value: value, Next: nil, Prev: nil}
	if d.Head == nil {
		d.Head = n
		d.Tail = n
	} else {
		lastNode := d.Tail
		lastNode.Next = n
		d.Tail = n
	}
	d.Length++
}

func (d *DoubleLinkedList) Remove(index int) {
	if !d.indexExist(index) {
		return
	}

	if index == 0 {
		d.Head = d.Head.Next
		d.Head.Prev = nil
		d.Length--
		return
	}
	prevNode := d.traverse(index - 1)
	if prevNode == nil {
		return
	}
	nextNode := prevNode.Next.Next
	prevNode.Next = nextNode
	nextNode.Prev = prevNode
	d.Length--
}

func (d *DoubleLinkedList) Update(index, value int) {
	n := d.traverse(index)
	n.Value = value
}

func (d *DoubleLinkedList) Lookup(index int) (*node, error) {
	if !d.indexExist(index) {
		return nil, errors.New(errorIndex)
	}
	return d.traverse(index), nil
}

func (d *DoubleLinkedList) Print() {
	n := d.Head
	for {
		if n == nil {
			break
		}
		if n != d.Tail {
			fmt.Print(n.Value, "-->")
		} else {
			fmt.Println(n.Value)
		}
		n = n.Next
	}
	fmt.Println(d.Length)
}
