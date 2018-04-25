package main

import "testing"

func TestP1(t *testing.T) {
	input, circuralList, expected := "3,4,1,5", []int{0, 1, 2, 3, 4}, 12
	if p1(input, circuralList) != expected {
		t.Errorf("p1(%v, %v) != %v", input, circuralList, expected)
	}
}

func TestP2(t *testing.T) {
	tables := []struct {
		input    string
		expected string
	}{
		{"", "a2582a3a0e66e6e86e3812dcb672a272"},
		{"AoC 2017", "33efeb34ea91902bb2f59c9920caa6cd"},
		{"1,2,3", "3efbe78a8d82f29979031a4aa0b16a9d"},
		{"1,2,4", "63960835bcdc130f0b66d7ff4f6a5a8e"}}

	for _, tab := range tables {
		if p2(tab.input) != tab.expected {
			t.Errorf("p2(%v) != %v", tab.input, tab.expected)
		}
	}
}
