package main

import "testing"

func TestP1(t *testing.T) {
	a, b := 65, 8921
	expected := 588
	if p1(a, b) != expected {
		t.Errorf("p1(%v, %v) != %v", a, b, expected)
	}
}

func TestP2(t *testing.T) {
	a, b := 65, 8921
	expected := 309
	if p2(a, b) != expected {
		t.Errorf("p2(%v, %v) != %v", a, b, expected)
	}
}
