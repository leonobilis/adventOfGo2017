package main

import "testing"

func TestP1(t *testing.T) {
	programs, input := "abcde", "s1,x3/4,pe/b"
	expected := "baedc"
	if p1(programs, input) != expected {
		t.Errorf("p1(programs, input)) != %v", expected)
	}
}

func TestP2(t *testing.T) {
	programs, input,  times:= "abcde", "s1,x3/4,pe/b", 2
	expected := "ceadb"
	if p2(programs, input, times) != expected {
		t.Errorf("p1(programs, input)) != %v", expected)
	}
}
