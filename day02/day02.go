package main

import (
	"adventOfGo2017/utils"
	"fmt"
	"strconv"
	"strings"
)

func p1(input string) int {
	sum := 0
	var min, max int
	for _, line := range strings.Split(input, "\n") {
		min = 1<<63 - 1
		max = 0
		for _, val := range strings.Fields(line) {
			v, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
		}
		sum += max - min
	}
	return sum
}

func p2(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		vals := make([]int, 0)
		for _, val := range strings.Fields(line) {
			v, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			vals = append(vals, v)
		}
		a, b := findDivisible(vals)

		sum += a / b
	}
	return sum
}

func findDivisible(ints []int) (int, int) {
	for _, a := range ints {
		for _, b := range ints {
			if a != b && a%b == 0 {
				return a, b
			}
		}
	}
	return 0, 0
}

func main() {
	input := utils.GetInput(2)
	fmt.Printf("Part 1: %v\n", p1(input))
	fmt.Printf("Part 2: %v\n", p2(input))
}
