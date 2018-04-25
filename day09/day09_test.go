package main

import "testing"

func TestP1(t *testing.T) {
	tables := []struct {
		input    string
		expected int
	}{
		{"{}", 1},
		{"{{{}}}", 6},
		{"{{},{}}", 5},
		{"{{{},{},{{}}}}", 16},
		{"{<a>,<a>,<a>,<a>}", 1},
		{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9},
		{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9},
		{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3}}

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
		{"<>", 0},
		{"<random characters>", 17},
		{"<<<<>", 3},
		{"<{!>}>", 2},
		{"<!!>", 0},
		{"<!!!>>", 0},
		{"<{o\"i!a,<{i<a>", 10}}

	for _, tab := range tables {
		if p2(tab.input) != tab.expected {
			t.Errorf("p2(%v) != %v", tab.input, tab.expected)
		}
	}
}
