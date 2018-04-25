package main

import "testing"

func TestP1(t *testing.T) {
	input := "0 2 7 0"
	expected := 5
	if p1(input) != expected {
		t.Errorf("p1(input) != %v", expected)
	}
}

func TestP2(t *testing.T) {
	input := "0 2 7 0"
	expected := 4
	if p2(input) != expected {
		t.Errorf("p2(input) != %v", expected)
	}
}
