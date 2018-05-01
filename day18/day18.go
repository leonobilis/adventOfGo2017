package main

import (
	"adventOfGo2017/utils"
	"fmt"
	"strconv"
	"strings"
)

func p1(input string) int {
	registers, instructions := make(map[byte]int), make([][]string, 0)
	val := func(arg string) int {
		if v, err := strconv.Atoi(arg); err == nil {
			return v
		}
		return registers[arg[0]]
	}
	for _, i := range strings.Split(input, "\n") {
		instructions = append(instructions, strings.Fields(i))
	}
	for pos, instruction, sndVal := 0, instructions[0], 0; pos < len(instructions); instruction = instructions[pos] {
		switch instruction[0] {
		case "snd":
			sndVal = registers[instruction[1][0]]
		case "set":
			registers[instruction[1][0]] = val(instruction[2])
		case "add":
			registers[instruction[1][0]] += val(instruction[2])
		case "mul":
			registers[instruction[1][0]] *= val(instruction[2])
		case "mod":
			registers[instruction[1][0]] %= val(instruction[2])
		case "rcv":
			if registers[instruction[1][0]] > 0 {
				return sndVal
			}
		case "jgz":
			if val(instruction[1]) > 0 {
				pos += val(instruction[2])
				continue
			}
		}
		pos++
	}
	return 0
}

func p2(input string) int {
	instructions := make([][]string, 0)
	for _, i := range strings.Split(input, "\n") {
		instructions = append(instructions, strings.Fields(i))
	}
	p1Chan, p2Chan, deadlock1Chan, deadlock2Chan, counterChan := make(chan int, 100), make(chan int, 100), make(chan bool, 1), make(chan bool, 1), make(chan int)
	proc := func(nr int, sndChan, rcvChan chan int, sndDeadlock, rcvDeadlock chan bool) {
		registers := make(map[byte]int)
		if nr == 1 {
			registers["p"[0]] = 1
		}
		val := func(arg string) int {
			if v, err := strconv.Atoi(arg); err == nil {
				return v
			}
			return registers[arg[0]]
		}

		for pos, instruction, counter := 0, instructions[0], 0; pos < len(instructions); instruction = instructions[pos] {
			switch instruction[0] {
			case "snd":
				select {
				case <-rcvDeadlock:
				default:
				}
				sndChan <- val(instruction[1])
				if nr == 1 {
					counter++
				}
			case "set":
				registers[instruction[1][0]] = val(instruction[2])
			case "add":
				registers[instruction[1][0]] += val(instruction[2])
			case "mul":
				registers[instruction[1][0]] *= val(instruction[2])
			case "mod":
				registers[instruction[1][0]] %= val(instruction[2])
			case "rcv":
				select {
				case registers[instruction[1][0]] = <-rcvChan:
				default:
					sndDeadlock <- true
					select {
					case registers[instruction[1][0]] = <-rcvChan:
					case <-rcvDeadlock:
						if nr == 1 {
							counterChan <- counter
						}
						return
					}
				}
			case "jgz":
				if val(instruction[1]) > 0 {
					pos += val(instruction[2])
					continue
				}
			}
			pos++
		}
	}

	go proc(0, p1Chan, p2Chan, deadlock1Chan, deadlock2Chan)
	go proc(1, p2Chan, p1Chan, deadlock2Chan, deadlock1Chan)

	return <-counterChan
}

func main() {
	input := utils.GetInput(18)
	fmt.Printf("Part 1: %v\n", p1(input))
	fmt.Printf("Part 2: %v\n", p2(input))
}
