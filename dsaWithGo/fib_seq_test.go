package main

import "testing"

func TestFibWRecTest(t *testing.T) {
	actual := fibWRec(6, nil)

	if actual != 8 {
		t.Errorf("expected 8, got %d", actual)
	}
}

func TestFibWoRecTest(t *testing.T) {
	actual := fibWoRec(6)

	if actual != 8 {
		t.Errorf("expected 8, got %d", actual)
	}
}
