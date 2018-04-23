package main

import "testing"

func TestP1(t *testing.T) {
	input := "5 1 9 5\n7 5 3\n2 4 6 8"
	expected := 18
	if p1(input) != expected {
		t.Errorf("p1(input) != %v", expected)
	}
}

func TestP2(t *testing.T) {
	input := "5 9 2 8\n9 4 7 3\n3 8 6 5"
	expected := 9
	if p2(input) != expected {
		t.Errorf("p2(input) != %v", expected)
	}
}
