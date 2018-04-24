package main

import "testing"

func TestP1(t *testing.T) {
	tables := []struct {
		input, expected int
	}{
		{1, 0},
		{12, 3},
		{23, 2},
		{1024, 31}}

	for _, tab := range tables {
		if p1(tab.input) != tab.expected {
			t.Errorf("p1(%v) != %v", tab.input, tab.expected)
		}
	}
}
