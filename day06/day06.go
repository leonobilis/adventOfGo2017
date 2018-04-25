package main

import (
	"adventOfGo2017/utils"
	"fmt"
	"reflect"
	"strings"
)

func parseInput(input string) []int {
	banks := make([]int, 0)
	for _, s := range strings.Fields(input) {
		val := utils.Atoi(s)
		banks = append(banks, val)
	}
	return banks
}

func seenBefore(banks []int, reached [][]int) bool {
	ch := make(chan bool)
	for _, r := range reached {
		go func(r []int) {
			if reflect.DeepEqual(banks, r) {
				ch <- true
			} else {
				ch <- false
			}
		}(r)
	}

	for i := 0; i < len(reached); i++ {
		if <-ch {
			return true
		}
	}

	return false
}

func maxIndex(banks []int) int {
	max, idx := 0, 0
	for i, v := range banks {
		if v > max {
			max = v
			idx = i
		}
	}
	return idx
}

func p1(input string) int {
	banks := parseInput(input)
	banksLen := len(banks)
	reached := make([][]int, 0)
	counter := 0
	for !seenBefore(banks, reached) {
		tmp := make([]int, banksLen)
		copy(tmp, banks)
		reached = append(reached, tmp)
		counter++
		toRedist := maxIndex(banks)
		redistVal := banks[toRedist]
		banks[toRedist] = 0
		for i := 1; i < redistVal+1; i++ {
			banks[(toRedist+i)%banksLen]++
		}
	}
	return counter
}

type reachedBank struct {
	bank []int
	when int
}

func seenBeforeRB(banks []int, reached []reachedBank) (bool, int) {
	type returnType struct {
		reached bool
		when    int
	}
	ch := make(chan returnType)
	for _, r := range reached {
		go func(r reachedBank) {

			if reflect.DeepEqual(banks, r.bank) {
				ch <- returnType{true, r.when}
			} else {
				ch <- returnType{false, 0}
			}
		}(r)
	}

	for i := 0; i < len(reached); i++ {
		if c := <-ch; c.reached {
			return true, c.when
		}
	}
	return false, 0
}

func p2(input string) int {
	banks := parseInput(input)
	banksLen := len(banks)
	reached := make([]reachedBank, 0)
	counter, whenSeenBefore := 0, 0
	for condition := false; !condition; condition, whenSeenBefore = seenBeforeRB(banks, reached) {
		tmp := make([]int, banksLen)
		copy(tmp, banks)
		reached = append(reached, reachedBank{tmp, counter})
		counter++
		toRedist := maxIndex(banks)
		redistVal := banks[toRedist]
		banks[toRedist] = 0
		for i := 1; i < redistVal+1; i++ {
			banks[(toRedist+i)%banksLen]++
		}
	}
	return counter - whenSeenBefore
}

func main() {
	input := utils.GetInput(6)
	fmt.Printf("Part 1: %v\n", p1(input))
	fmt.Printf("Part 2: %v\n", p2(input))
}
