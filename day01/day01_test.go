package main

import "testing"

func TestP1(t *testing.T) {
	tables := []struct {
		input    string
		expected int
	}{
		{"1122", 3},
		{"1111", 4},
		{"1234", 0},
		{"91212129", 9}}

	for _, tab := range tables {
		if p1(tab.input) != tab.expected {
			t.Errorf("p1(%v) != %v", tab.input, tab.expected)
		}
	}
}

func TestP2(t *testing.T) {
	tables := []struct {
		input    string
		expected int
	}{
		{"1212", 6},
		{"1221", 0},
		{"123425", 4},
		{"123123", 12},
		{"12131415", 4}}

	for _, tab := range tables {
		if p2(tab.input) != tab.expected {
			t.Errorf("p2(%v) != %v", tab.input, tab.expected)
		}
	}
}
