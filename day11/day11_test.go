package main

import "testing"

func TestP1(t *testing.T) {
	tables := []struct {
		input    string
		expected int
	}{
		{"ne,ne,ne", 3},
		{"ne,ne,sw,sw", 0},
		{"ne,ne,s,s", 2},
		{"se,sw,se,sw,sw", 3}}

	for _, tab := range tables {
		if p1(tab.input) != tab.expected {
			t.Errorf("p1(%v) != %v", tab.input, tab.expected)
		}
	}
}
