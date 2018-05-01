package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"fmt"
)

func getInput(i int) string {
	input, e := ioutil.ReadFile(filepath.Join(os.Getenv("GOPATH"), fmt.Sprintf("src/adventOfGo2017/day18/test_input%d.txt", i)))
	if e != nil {
		panic(e)
	}
	return string(input)
}

func TestP1(t *testing.T) {
	input := getInput(1)
	expected := 4
	if p1(input) != expected {
		t.Errorf("p1(input) != %v", expected)
	}
}

func TestP2(t *testing.T) {
	input := getInput(2)
	expected := 3
	if p2(input) != expected {
		t.Errorf("p2(input) != %v", expected)
	}
}
