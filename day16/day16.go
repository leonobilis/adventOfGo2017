package main

import (
	"adventOfGo2017/utils"
	"fmt"
	"strings"
)

func spin(inlist, x string) string {
	return inlist[len(inlist)-utils.Atoi(x):] + inlist[:len(inlist)-utils.Atoi(x)]
}

func exchange(inlist, ab string) string {
	splitted := strings.Split(ab, "/")
	a, b := utils.Atoi(splitted[0]), utils.Atoi(splitted[1])
	if a > b {
		a, b = b, a
	}
	return inlist[:a] + inlist[b:b+1] + inlist[a+1:b] + inlist[a:a+1] + inlist[b+1:]
}

func partner(inlist, ab string) string {
	splitted := strings.Split(ab, "/")
	a, b := splitted[0], splitted[1]
	ia, ib := strings.Index(inlist, splitted[0]), strings.Index(inlist, splitted[1])
	if ia > ib {
		ia, ib = ib, ia
		a, b = b, a
	}
	return inlist[:ia] + b + inlist[ia+1:ib] + a + inlist[ib+1:]
}

func dance(programs, input string) string {
	action := map[byte]func(string, string) string{
		"s"[0]: spin,
		"x"[0]: exchange,
		"p"[0]: partner,
	}
	for _, i := range strings.Split(input, ",") {
		programs = action[i[0]](programs, i[1:])
	}
	return programs
}

func p1(programs, input string) string {
	return dance(programs, input)
}

func p2(programs, input string, times int) string {
	programsCpy, count := programs, 1
	for programs = dance(programs, input); programs != programsCpy; programs = dance(programs, input) {
		count++
	}
	for i := 0; i < times%count; i++ {
		programs = dance(programs, input)
	}
	return programs
}

func main() {
	input := utils.GetInput(16)
	fmt.Printf("Part 1: %v\n", p1("abcdefghijklmnop", input))
	fmt.Printf("Part 2: %v\n", p2("abcdefghijklmnop", input, 1000000000))
}
