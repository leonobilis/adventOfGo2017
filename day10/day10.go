package main

import (
	"adventOfGo2017/utils"
	"fmt"
	"strings"
)

func p1(input string, circuralList []int) int {
	length, p1, skip := len(circuralList), 0, 0
	for _, i := range strings.Split(input, ",") {
		ii := utils.Atoi(i)
		p2 := (p1 + ii) % length
		if p2 > p1 {
			circuralList = append(circuralList[:p1], append(utils.Reverse(circuralList[p1:p2]), circuralList[p2:]...)...)
		} else if ii > 0 {
			sublist := utils.Reverse(append(circuralList[p1:], circuralList[:p2]...))
			circuralList = append(sublist[length-p1:], append(circuralList[p2:p1], sublist[:length-p1]...)...)
		}
		p1 = (p1 + ii + skip) % length
		skip++
	}
	return circuralList[0] * circuralList[1]
}

func p2(input string) string {
	circuralList := utils.MakeRange(0, 255)
	lengths := make([]int, 0)
	for _, i := range input {
		lengths = append(lengths, int(i))
	}
	lengths = append(lengths, []int{17, 31, 73, 47, 23}...)

	length, p1, skip := len(circuralList), 0, 0
	for i := 0; i < 64; i++ {
		for _, l := range lengths {
			p2 := (p1 + l) % length
			if p2 > p1 {
				circuralList = append(circuralList[:p1], append(utils.Reverse(circuralList[p1:p2]), circuralList[p2:]...)...)
			} else if l != 0 {
				sublist := utils.Reverse(append(circuralList[p1:], circuralList[:p2]...))
				circuralList = append(sublist[length-p1:], append(circuralList[p2:p1], sublist[:length-p1]...)...)
			}
			p1 = (p1 + l + skip) % length
			skip++
		}
	}

	hash := ""
	for i := 0; i < 256; i += 16 {
		xor := 0
		for _, c := range circuralList[i : i+16] {
			xor ^= c
		}
		hash = fmt.Sprintf("%s%02x", hash, xor)
	}

	return hash
}

func main() {
	input := utils.GetInput(10)
	fmt.Printf("Part 1: %v\n", p1(input, utils.MakeRange(0, 255)))
	fmt.Printf("Part 2: %v\n", p2(input))
}
