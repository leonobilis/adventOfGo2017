package main

import (
	"fmt"
	"adventOfGo2017/utils"
)

func p1(input int) int {
	buffer, pos := []int{0}, 0
	for _, i := range utils.MakeRange(1, 2017) {
		pos = (pos+1+input) % i
		buffer = append(buffer[:pos+1], append( []int{i}, buffer[pos+1:]...)...)
	}
	return buffer[pos+2]
}

func p2(input int) int {
	pos, current := 0, 0
	for _, i := range utils.MakeRange(1, 50000000) {
		pos = (pos+1+input) % i
		if pos == 0 {
			current = i
		}
	}
	return current
}

func main() {
	input := 348
	fmt.Printf("Part 1: %v\n", p1(input))
	fmt.Printf("Part 2: %v\n", p2(input))
}
