package main

import "testing"

func TestP1(t *testing.T) {
	input := 3
	expected := 638
	if p1(input) != expected {
		t.Errorf("p1(%v) != %v", input, expected)
	}
}
