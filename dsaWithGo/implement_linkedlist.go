package main

import "fmt"

type Move struct {
	// the value of the move
	data string
	next *Move // pointer next element of exact same type
}

type ChessMatch struct {
	head *Move // pointer to the first element of the linked list
}

// This initializes a new ChessMatch struct
func NewChessMatch() *ChessMatch {
	return &ChessMatch{}
	// returns address of the ChessMatch struct
}

// This facilitates Move is populated with data and
// the next element in the linked list
func NewMove(data string, next *Move) *Move {
	return &Move{data: data, next: next}
}

func (c *ChessMatch) getAt(index int) *Move {
	crt := c.head // current value of head

	pos := 0 // current position of iterator

	for pos < index && crt != nil {
		crt = crt.next
		pos++
	}

	// when desired index is reached,
	// return the current head value
	return crt
}

func (c *ChessMatch) getLast() *Move {
	if c.head == nil {
		return nil
	}

	current := c.head

	// traverse the linked list until the last element
	for current != nil && current.next != nil {
		current = current.next
	}

	return current
}

func (c *ChessMatch) insertAt(index int, data string) {
	if c.head == nil {
		c.head = NewMove(data, nil)
		return
	}

	// inserting at the beginning of the list
	if index == 0 {

		// save the current head value
		h := c.head

		// set the new head value to the new element
		c.head = NewMove(data, h)
		return
	}

	// before inserting at position 5 we need to modify
	// next prop of element at position 4 to point to the new element
	// the lower logic didn't handle, handling element at head,
	// because it relied on previous element
	// to be able to insert the new element
	prev := c.getAt(index - 1)

	// if the element at the desired index is not found
	// insert at the end of the list
	if prev == nil {
		prev = c.getLast()
	}
	// here we are modifying next prop of element at position 5
	// to point to next of element at position 4 before new
	// element is inserted
	prev.next = NewMove(data, prev.next)

}

func (c *ChessMatch) removeAt(index int) {
	if c.head == nil {
		return
	}

	// remove the head element

	if index == 0 {
		c.head = c.head.next
		return
	}

	prev := c.getAt(index - 1)

	if prev != nil && prev.next != nil {
		prev.next = prev.next.next
	}
}

func (c *ChessMatch) forEach(predicatefn func(*Move)) {
	if c.head == nil {
		return
	}

	current := c.head

	for current != nil {

		// call the predicate function with the current element
		predicatefn(current)
		current = current.next
	}
}

// example use case of predicate function

func (c *ChessMatch) replayMatch() {
	printMove := func(t *Move) {
		// print the current element
		fmt.Println((t.data))
	}
	c.forEach(printMove)
}
