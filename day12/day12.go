package main

import (
	"adventOfGo2017/utils"
	"fmt"
	"regexp"
	"strings"
)

type programs map[int][]int

type set map[int]struct{}

func (s set) contains(key int) bool {
	_, ok := s[key]
	return ok
}

func (s set) pop() int {
	for pop := range s {
		delete(s, pop)
		return pop
	}
	return 0
}
func (s set) append(key int) {
	s[key] = struct{}{}
}

func parseInput(input string) programs {
	prog := make(programs)
	re := regexp.MustCompile("([\\d]+) <-> ([\\d, ]+)")
	for _, s := range strings.Split(input, "\n") {
		groups := re.FindStringSubmatch(s)
		number := utils.Atoi(groups[1])
		pipes := make([]int, 0)
		for _, p := range strings.Split(groups[2], ", ") {
			pipes = append(pipes, utils.Atoi(p))
		}
		prog[number] = pipes
	}
	return prog
}

func p1(input string) int {
	prog := parseInput(input)
	keys, toProcess := make(set), set{0: {}}
	for len(toProcess) > 0 {
		for _, k := range prog[toProcess.pop()] {
			if !toProcess.contains(k) && !keys.contains(k) {
				toProcess.append(k)
				keys.append(k)
			}
		}
	}
	return len(keys)
}

func recursiveDelete(p programs, prog int) {
	if pipes, ok := p[prog]; ok {
		delete(p, prog)
		for _, i := range pipes {
			recursiveDelete(p, i)
		}
	}
}

func p2(input string) int {
	prog := parseInput(input)
	i := 0
	for len(prog) > 0 {
		for p := range prog {
			recursiveDelete(prog, p)
			break
		}
		i++
	}
	return i
}

func main() {
	input := utils.GetInput(12)
	fmt.Printf("Part 1: %v\n", p1(input))
	fmt.Printf("Part 2: %v\n", p2(input))
}
