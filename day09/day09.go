package main

import (
	"adventOfGo2017/utils"
	"fmt"
)

func p1(input string) int {
	score, opened, group_score, garbage, ignore := 0, make([]int, 0), 1, false, false
	for _, i := range input {
		switch {
		case ignore:
			ignore = false
		case i == rune("!"[0]):
			ignore = true
		case garbage:
			if i == rune(">"[0]) {
				garbage = false
			}
		default:
			switch i {
			case rune("<"[0]):
				garbage = true
			case rune("{"[0]):
				opened = append(opened, group_score)
				group_score++
			case rune("}"[0]):
				score, opened = score+opened[len(opened)-1], opened[:len(opened)-1]
				group_score--
			}
		}
	}
	return score
}

func p2(input string) int {
	garbage_counter, garbage, ignore := 0, false, false
	for _, i := range input {
		switch {
		case ignore:
			ignore = false
		case i == rune("!"[0]):
			ignore = true
		case garbage:
			if i == rune(">"[0]) {
				garbage = false
			} else {
				garbage_counter++
			}
		default:
			if i == rune("<"[0]) {
				garbage = true
			}
		}
	}
	return garbage_counter
}

func main() {
	input := utils.GetInput(9)
	fmt.Printf("Part 1: %v\n", p1(input))
	fmt.Printf("Part 2: %v\n", p2(input))
}
