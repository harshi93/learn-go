package main

import "testing"

func TestGetAtHeadPosition(t *testing.T) {
	c := NewChessMatch()

	mData := "e4"
	m2Data := "Nf6"

	c.insertAt(0, mData)
	c.insertAt(1, m2Data)

	actual := c.getAt(0)

	if actual == nil || actual.data != mData {
		t.Errorf("Expected %s, got %s", mData, actual.data)
	}
}

func TestGetAtOutOfBounds(t *testing.T) {
	c := NewChessMatch()

	mData := "e4"

	c.insertAt(0, mData)

	actual := c.getAt(10)

	if actual != nil {
		t.Errorf("Expected nil, got %s", actual.data)
	}
}

func TestGetAtLastPosition(t *testing.T) {
	c := NewChessMatch()

	mData := "e4"
	m2Data := "Nf6"

	c.insertAt(0, mData)
	c.insertAt(1, m2Data)

	actual := c.getAt(1)

	if actual == nil || actual.data != m2Data {
		t.Errorf("Expected %s, got %s", m2Data, actual.data)
	}
}

func TestInsertAtHeadPosition(t *testing.T) {
	c := NewChessMatch()

	mData := "e4"
	m2Data := "e5"

	c.insertAt(0, mData)
	c.insertAt(0, m2Data)

	if c.head == nil || c.head.data != m2Data {
		t.Errorf("Expected %s, got %s", m2Data, c.head.data)
	}
}

func TestInsertAtRandomPosition(t *testing.T) {
	c := NewChessMatch()

	mData := "e4"
	m2Data := "Nf6"
	m3Data := "Nc3"
	m4Data := "Bb5"

	c.insertAt(0, mData)
	c.insertAt(1, m2Data)
	c.insertAt(2, m3Data)
	c.insertAt(3, m4Data)

	actual := c.getAt(2)

	if actual == nil || actual.data != m3Data {
		t.Errorf("Expected %s, got %s", m3Data, actual.data)
	}
}

func TestAtInsertAtOutOfBounds(t *testing.T) {
	c := NewChessMatch()

	mData := "e4"
	m2Data := "Nf6"

	c.insertAt(0, mData)
	c.insertAt(10, m2Data)

	if c.getAt(1) == nil || c.getAt(1).data != m2Data {
		t.Errorf("Expected element %s at position 1", m2Data)
	}
}

func TestRemoveAtHeadPosition(t *testing.T) {
	c := NewChessMatch()

	mData := "e4"
	m2Data := "Nf6"

	c.insertAt(0, mData)
	c.insertAt(1, m2Data)

	c.removeAt(0)

	actual := c.getAt(0)

	if actual == nil || actual.data != m2Data {
		t.Errorf("Expected %s, got %s", m2Data, c.head.data)
	}
}

func TestRemoveAtRandomPosition(t *testing.T) {
	c := NewChessMatch()

	mData := "e4"
	m2Data := "Nf6"
	m3Data := "Nc3"

	c.insertAt(0, mData)
	c.insertAt(1, m2Data)
	c.insertAt(2, m3Data)

	c.removeAt(1)

	actual := c.getAt(1)

	if actual == nil || actual.data != m3Data {
		t.Errorf("Expected %s, got %s", m2Data, actual.data)
	}
}

func TestRemoveAtLastPosition(t *testing.T) {
	c := NewChessMatch()

	mData := "e4"
	m2Data := "Nf6"

	c.insertAt(0, mData)
	c.insertAt(1, m2Data)

	c.removeAt(1)

	actualLast := c.getLast()

	if actualLast == nil || actualLast.data != mData {
		t.Errorf("Expected %s, got %s", mData, actualLast.data)
	}

}

func TestRemoveAtOutOfBounds(t *testing.T) {
	c := NewChessMatch()

	mData := "e4"
	m2Data := "Nf6"

	c.insertAt(0, mData)
	c.insertAt(1, m2Data)

	c.removeAt(10)

	// this test should pass if everything works
	// as expected otherwise go will panic
}

func TestForEach(t *testing.T) {
	c := NewChessMatch()

	mData := "e4"
	m2Data := "Nf6"
	m3Data := "Nc3"
	m4Data := "Bb5"

	c.insertAt(0, mData)
	c.insertAt(1, m2Data)
	c.insertAt(2, m3Data)
	c.insertAt(3, m4Data)

	newData := "modified data"

	modifyDataFn := func(element *Move) {
		element.data = newData
	}

	c.forEach(modifyDataFn)

	actual1 := c.getAt(0)
	actual2 := c.getAt(1)
	actual3 := c.getAt(2)
	actual4 := c.getAt(3)

	if (actual1 == nil || actual1.data != newData) ||
		(actual2 == nil || actual2.data != newData) ||
		(actual3 == nil || actual3.data != newData) ||
		(actual4 == nil || actual4.data != newData) {
		t.Errorf("Expected %s, got %s", newData, actual1.data)
	}
}
