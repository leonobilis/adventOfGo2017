package main

import (
	"adventOfGo2017/utils"
	"fmt"
	"sort"
	"strings"
)

func check1(input string) bool {
	words := strings.Fields(input)
	sort.Strings(words)
	for i := 0; i < len(words)-1; i++ {
		if words[i] == words[i+1] {
			return false
		}
	}
	return true
}

func check2(input string) bool {
	words := strings.Fields(input)
	for i, w := range words {
		words[i] = utils.SortString(w)
	}
	sort.Strings(words)
	for i := 0; i < len(words)-1; i++ {
		if words[i] == words[i+1] {
			return false
		}
	}
	return true
}

func part(input string, check func(string) bool) int {
	passphrases := strings.Split(input, "\n")
	valid := make(chan bool)
	sum := 0
	for _, p := range passphrases {
		go func(pass string) {
			valid <- check(pass)
		}(p)
	}
	for i := 0; i < len(passphrases); i++ {
		if <-valid {
			sum++
		}
	}
	return sum
}

func p1(input string) int {
	return part(input, check1)
}

func p2(input string) int {
	return part(input, check2)
}

func main() {
	input := utils.GetInput(4)
	fmt.Printf("Part 1: %v\n", p1(input))
	fmt.Printf("Part 2: %v\n", p2(input))
}
