package main

import (
	"adventOfGo2017/utils"
	"fmt"
	"strings"
)

type hexTiling struct {
	x, y, z int
}

func (h *hexTiling) Move(dir string) {
	switch dir {
	case "n":
		h.x++
		h.y++
	case "ne":
		h.y++
		h.z++
	case "se":
		h.z++
		h.x--
	case "s":
		h.x--
		h.y--
	case "sw":
		h.y--
		h.z--
	case "nw":
		h.z--
		h.x++
	}
}

func (h hexTiling) FromStart() int {
	return (utils.Abs(h.x) + utils.Abs(h.y) + utils.Abs(h.z)) / 2
}

func p1(input string) int {
	var h hexTiling
	for _, dir := range strings.Split(input, ",") {
		h.Move(dir)
	}
	return h.FromStart()
}

func p2(input string) int {
	var h hexTiling
	max := 0
	for _, dir := range strings.Split(input, ",") {
		h.Move(dir)
		fromStart := h.FromStart()
		if fromStart > max {
			max = fromStart
		}
	}
	return max
}

func main() {
	input := utils.GetInput(11)
	fmt.Printf("Part 1: %v\n", p1(input))
	fmt.Printf("Part 2: %v\n", p2(input))
}
