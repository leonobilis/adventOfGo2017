package main

import (
	"adventOfGo2017/utils"
	"fmt"
	"strings"
)

func parseInput(input string) []int {
	offsets := make([]int, 0)
	for _, s := range strings.Fields(input) {
		val := utils.Atoi(s)
		offsets = append(offsets, val)
	}
	return offsets
}

func p1(input string) int {
	offsets := parseInput(input)
	pos, step := 0, 0
	for 0 <= pos && pos < len(offsets) {
		tmp := offsets[pos]
		offsets[pos]++
		pos += tmp
		step++
	}
	return step
}

func p2(input string) int {
	offsets := parseInput(input)
	pos, step := 0, 0
	for 0 <= pos && pos < len(offsets) {
		tmp := offsets[pos]
		if tmp >= 3 {
			offsets[pos]--
		} else {
			offsets[pos]++
		}
		pos += tmp
		step++
	}
	return step
}

func main() {
	input := utils.GetInput(5)
	fmt.Printf("Part 1: %v\n", p1(input))
	fmt.Printf("Part 2: %v\n", p2(input))
}
