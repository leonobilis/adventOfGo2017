package main

import (
	"adventOfGo2017/utils"
	"fmt"
	"regexp"
	"strings"
)

type scanners map[int]int

func parseInput(input string) scanners {
	sc := make(scanners)
	re := regexp.MustCompile("([\\d]+): ([\\d]+)")
	for _, s := range strings.Split(input, "\n") {
		groups := re.FindStringSubmatch(s)
		sc[utils.Atoi(groups[1])] = utils.Atoi(groups[2])
	}
	return sc
}

func p1(input string) int {
	sc, sum := parseInput(input), 0
	for depth, rang := range sc {
		if a := depth % ((rang - 1) * 2); rang > 0 && a == 0 || ((rang-1)*2-a) == 0 {
			sum += depth * rang
		}
	}
	return sum
}

func p2(input string) int {
	sc, delay := parseInput(input), 0

	for cond := true; cond; {
		for depth, rang := range sc {
			if a := (depth + delay) % ((rang - 1) * 2); rang > 0 && a == 0 || ((rang-1)*2-a) == 0 {
				cond = true
				delay++
				break
			}
			cond = false
		}
	}
	return delay
}

func main() {
	input := utils.GetInput(13)
	fmt.Printf("Part 1: %v\n", p1(input))
	fmt.Printf("Part 2: %v\n", p2(input))
}
