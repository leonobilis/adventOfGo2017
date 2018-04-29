package main

import "testing"

func TestP1(t *testing.T) {
	input := "flqrgnkx"
	expected := 8108
	if p1(input) != expected {
		t.Errorf("p1(input) != %v", expected)
	}
}

func TestP2(t *testing.T) {
	input := "flqrgnkx"
	expected := 1242
	if p2(input) != expected {
		t.Errorf("p2(input) != %v", expected)
	}
}
