package main

import (
	"adventOfGo2017/utils"
	"fmt"
)

func p1(input string) int {
	length := len(input)
	sum := 0
	for i, val := range input {
		if val == rune(input[(i+1)%length]) {
			sum += int(val - '0')
		}
	}
	return sum
}

func p2(input string) int {
	length := len(input)
	half := length / 2
	sum := 0
	for i, val := range input {
		if val == rune(input[(i+half)%length]) {
			sum += int(val - '0')
		}
	}
	return sum
}

func main() {
	input := utils.GetInput(1)
	fmt.Printf("Part 1: %v\n", p1(input))
	fmt.Printf("Part 2: %v\n", p2(input))
}
