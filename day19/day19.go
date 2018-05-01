package main

import (
	"adventOfGo2017/utils"
	"fmt"
	"strings"
	"unicode"
)

type pair struct {
	y, x int
}

var (
	down  = pair{1, 0}
	up    = pair{-1, 0}
	right = pair{0, 1}
	left  = pair{0, -1}
)

func parseInput(input string) []string {
	out := make([]string, 0)
	for _, s := range strings.Split(input, "\n") {
		out = append(out, strings.Replace(s, " ", ".", -1))
	}
	return out
}

func route(input string) (letters string, count int) {
	inp := parseInput(input)
	for direction, x, y, end := down, strings.Index(inp[0], "|"), 0, false; !end; x, y, count = x+direction.x, y+direction.y, count+1 {
		s := inp[y][x]
		switch {
		case s == "-"[0] || s == "|"[0]:
		case unicode.IsLetter(rune(s)):
			letters += string(s)
		case s == "+"[0]:
			switch {
			case direction != up && y+1 < len(inp) && (inp[y+1][x] == "|"[0] || unicode.IsLetter(rune(inp[y+1][x]))):
				direction = down
			case direction != down && y-1 >= 0 && (inp[y-1][x] == "|"[0] || unicode.IsLetter(rune(inp[y-1][x]))):
				direction = up
			case direction != left && x+1 < len(inp[y]) && (inp[y][x+1] == "-"[0] || unicode.IsLetter(rune(inp[y][x+1]))):
				direction = right
			case direction != right && x-1 >= 0 && (inp[y][x-1] == "-"[0] || unicode.IsLetter(rune(inp[y][x-1]))):
				direction = left
			default:
				end = true
			}
		default:
			end = true
		}
	}
	return letters, count-1
}

func p1(input string) string {
	letters , _ := route(input)
	return letters

}

func p2(input string) int {
	_ , count := route(input)
	return count
}

func main() {
	input := utils.GetInput(19)
	fmt.Printf("Part 1: %v\n", p1(input))
	fmt.Printf("Part 2: %v\n", p2(input))
}
