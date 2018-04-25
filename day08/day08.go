package main

import (
	"adventOfGo2017/utils"
	"fmt"
	"strings"
)

type action func(int) int
type condition func(int) bool

type instruction struct {
	regAct  string
	act     action
	regCond string
	cond    condition
}

func parseInput(input string) []instruction {
	actions := map[string]func(int) action{
		"inc": func(value int) action { return func(reg int) int { return reg + value } },
		"dec": func(value int) action { return func(reg int) int { return reg - value } }}
	conditions := map[string]func(int) condition{
		">":  func(value int) condition { return func(reg int) bool { return reg > value } },
		"<":  func(value int) condition { return func(reg int) bool { return reg < value } },
		">=": func(value int) condition { return func(reg int) bool { return reg >= value } },
		"<=": func(value int) condition { return func(reg int) bool { return reg <= value } },
		"==": func(value int) condition { return func(reg int) bool { return reg == value } },
		"!=": func(value int) condition { return func(reg int) bool { return reg != value } }}
	instructions := make([]instruction, 0)

	for _, line := range strings.Split(input, "\n") {
		l := strings.Fields(line)
		instructions = append(instructions, instruction{
			regAct:  l[0],
			act:     actions[l[1]](utils.Atoi(l[2])),
			regCond: l[4],
			cond:    conditions[l[5]](utils.Atoi(l[6]))})
	}
	return instructions
}

func p1(input string) int {
	instructions := parseInput(input)
	registers := make(map[string]int)
	for _, i := range instructions {
		if i.cond(registers[i.regCond]) {
			registers[i.regAct] = i.act(registers[i.regAct])
		}
	}
	max := -1 << 31
	for _, v := range registers {
		if v > max {
			max = v
		}
	}
	return max
}

func p2(input string) int {
	instructions := parseInput(input)
	registers := make(map[string]int)
	max := -1 << 31
	for _, i := range instructions {
		if i.cond(registers[i.regCond]) {
			registers[i.regAct] = i.act(registers[i.regAct])
			if registers[i.regAct] > max {
				max = registers[i.regAct]
			}
		}
	}
	return max
}

func main() {
	input := utils.GetInput(8)
	fmt.Printf("Part 1: %v\n", p1(input))
	fmt.Printf("Part 2: %v\n", p2(input))
}
