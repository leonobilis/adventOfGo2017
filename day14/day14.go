package main

import (
	"adventOfGo2017/utils"
	"fmt"
	"strconv"
	"strings"
)

func hash(input string) string {
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
		hash = fmt.Sprintf("%s%08b", hash, xor)
	}

	return hash
}

func p1(input string) int {
	sum := 0
	for i := 0; i < 128; i++ {
		sum += strings.Count(hash(fmt.Sprintf("%v-%v", input, strconv.Itoa(i))), "1")
	}
	return sum
}

func proc(grid [][128]byte, i, j int) bool {
	if i >= 0 && i < 128 && j >= 0 && j < 128 && grid[i][j] == "1"[0] {
		grid[i][j] = "0"[0]
		proc(grid, i+1, j)
		proc(grid, i-1, j)
		proc(grid, i, j+1)
		proc(grid, i, j-1)
		return true
	}
	return false
}

func p2(input string) int {
	grid := make([][128]byte, 128)
	for i := 0; i < 128; i++ {
		reader := strings.NewReader(hash(fmt.Sprintf("%v-%v", input, strconv.Itoa(i))))
		reader.Read(grid[i][:128])
	}
	sum := 0
	for i := 0; i < 128; i++ {
		for j := 0; j < 128; j++ {
			if proc(grid, i, j) {
				sum++
			}
		}
	}
	return sum
}

func main() {
	input := utils.GetInput(14)
	fmt.Printf("Part 1: %v\n", p1(input))
	fmt.Printf("Part 2: %v\n", p2(input))
}
