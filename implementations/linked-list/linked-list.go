package main

import (
	"errors"
	"fmt"
)

type Element struct {
	value int
	next  *Element
}

type LinkedList struct {
	head *Element
	len  int
}

func (l *LinkedList) InsertFirst(v int) {
	var curr = Element{}
	curr.value = v
	if curr.next != nil {
		curr.next = l.head
	}
	l.head = &curr
	l.len++
}

// IterateList method iterates over LinkedList
func (l *LinkedList) IterateList() {
	var e *Element
	for e = l.head; e != nil; e = e.next {
		fmt.Println(e.value)
	}
}

// Last method returns the last Element
func (l *LinkedList) Last() *Element {
	var e *Element
	for e = l.head; e != nil; e = e.next {
		if e.next == nil {
			return e
		}
	}
	return nil
}

// InsertLast method adds the Element with value to the end
func (l *LinkedList) InsertLast(v int) {
	var e = &Element{}
	e.value = v
	e.next = nil
	var lastElem = l.Last()
	if lastElem != nil {
		lastElem.next = e
	}
}

//SearchByValue method returns Element given parameter value
func (l *LinkedList) SearchByValue(v int) *Element {
	var e = &Element{}
	for e = l.head; e.next != nil; e = e.next {
		if e.value == v {
			return e
		}
	}
	return nil
}

// AddAfter method adds a node with nodeProperty after node with property
func (l *LinkedList) AddAfter(after int, v int) {
	var e = &Element{}
	e.value = v
	e.next = nil
	var nodeWith *Element
	nodeWith = l.SearchByValue(after)
	if nodeWith != nil {
		e.next = nodeWith.next
		nodeWith.next = e
	}
}

func (l *LinkedList) Delete(v int) error {
	var e *Element
	if l.head.value == v {
		l.head = l.head.next
	}
	for e = l.head; e.next != nil; e = e.next {
		if e.next.value == v {
			e.next = e.next.next
			return nil
		}
		// for last element
		if e.next.next == nil {
			e.next = nil
		}
	}
	return errors.New("VALUE NOT FOUND")
}

func main() {
	var ll LinkedList
	ll = LinkedList{}
	ll.InsertFirst(1)
	ll.InsertLast(3)
	ll.InsertLast(5)
	ll.AddAfter(1, 7)
	//ll.IterateList()
	var _ = ll.Delete(5)
	ll.IterateList()
}
