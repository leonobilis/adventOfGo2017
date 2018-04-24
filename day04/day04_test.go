package main

import "testing"

func TestP1(t *testing.T) {
	tables := []struct {
		input    string
		expected bool
	}{
		{"aa bb cc dd ee", true},
		{"aa bb cc dd aa", false},
		{"aa bb cc dd aaa", true}}

	for _, tab := range tables {
		if check1(tab.input) != tab.expected {
			t.Errorf("check1(%v) != %v", tab.input, tab.expected)
		}
	}
}

func TestP2(t *testing.T) {
	tables := []struct {
		input    string
		expected bool
	}{
		{"abcde fghij", true},
		{"abcde xyz ecdab", false},
		{"a ab abc abd abf abj", true},
		{"iiii oiii ooii oooi oooo", true},
		{"oiii ioii iioi iiio", false}}

	for _, tab := range tables {
		if check2(tab.input) != tab.expected {
			t.Errorf("check2(%v) != %v", tab.input, tab.expected)
		}
	}
}
